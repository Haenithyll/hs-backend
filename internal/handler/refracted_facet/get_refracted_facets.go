package refracted_facet

import (
	dto "hs-backend/internal/dto/refracted_facet"
	"hs-backend/internal/handler"
	"hs-backend/internal/model"
	"hs-backend/internal/repository"
	"hs-backend/internal/types"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetRefractedFacetsHandler struct {
	FacetRepository            repository.FacetRepository
	UserRepository             repository.UserRepository
	UserPrismTrackerRepository repository.UserPrismTrackerRepository
}

func NewGetRefractedFacetsHandler(
	facetRepository repository.FacetRepository,
	userRepository repository.UserRepository,
	userPrismTrackerRepository repository.UserPrismTrackerRepository,
) *GetRefractedFacetsHandler {
	return &GetRefractedFacetsHandler{
		facetRepository,
		userRepository,
		userPrismTrackerRepository,
	}
}

// GetRefractedFacetsHandler godoc
// @Summary Get refracted facets
// @Description Returns refracted facets
// @Tags Refracted Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetRefractedFacetsResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/refracted-facets [get]
func (h *GetRefractedFacetsHandler) Handle(c *gin.Context) {
	facetRepository := h.FacetRepository
	userRepository := h.UserRepository
	userPrismTrackerRepository := h.UserPrismTrackerRepository

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	users, err := userRepository.FindAll()
	if err != nil {
		handler.InternalError(c, "Failed to get users: "+err.Error())
		return
	}

	userPrismTrackers, err := userPrismTrackerRepository.FindAllWithPrisms()
	if err != nil {
		handler.InternalError(c, "Failed to get user prism trackers: "+err.Error())
		return
	}

	lastUpdatedAtByFacetIdMap := h.getLastUpdatedAtByFacetIdMap(userPrismTrackers, userId)

	if len(lastUpdatedAtByFacetIdMap) == 0 {
		handler.OK(c, dto.ToGetRefractedFacetsResponse(make(map[uuid.UUID]types.RefractedFacet), users, userId))
		return
	}

	facetIds := make([]uint8, 0, len(lastUpdatedAtByFacetIdMap))
	for facetId := range lastUpdatedAtByFacetIdMap {
		facetIds = append(facetIds, facetId)
	}

	facets, err := facetRepository.FindManyRefractedByIds(facetIds)
	if err != nil {
		handler.InternalError(c, "Failed to get facets: "+err.Error())
		return
	}

	refractedFacetByUserIdMap := make(map[uuid.UUID]types.RefractedFacet)

	for _, facet := range facets {
		refractedFacetByUserIdMap[facet.User.ID] = types.RefractedFacet{
			Id:            facet.ID,
			Label:         facet.PublicLabel,
			Color:         facet.Color,
			LastUpdatedAt: lastUpdatedAtByFacetIdMap[facet.ID],
		}
	}

	response := dto.ToGetRefractedFacetsResponse(refractedFacetByUserIdMap, users, userId)
	handler.OK(c, response)
}

func (h *GetRefractedFacetsHandler) getLastUpdatedAtByFacetIdMap(
	userPrismTrackers []model.UserPrismTracker,
	userId uuid.UUID,
) map[uint8]time.Time {
	lastUpdatedAtByFacetIdMap := make(map[uint8]time.Time)

	for _, upt := range userPrismTrackers {
		if upt.UserId == userId {
			continue
		}

		config := upt.Prism.Configuration

		var facetId uint8 = config.Base
		var lastUpdatedAt time.Time = upt.LastUpdatedAt

		for _, configUserItems := range config.Users {
			if configUserItems.UserId == userId {
				facetId = configUserItems.FacetId
				lastUpdatedAt = upt.LastUpdatedAt
				break
			}
		}

		lastUpdatedAtByFacetIdMap[facetId] = lastUpdatedAt
	}

	return lastUpdatedAtByFacetIdMap
}
