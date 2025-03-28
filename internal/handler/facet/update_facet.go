package facet

import (
	"fmt"
	dto "hs-backend/internal/dto/facet"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateFacetHandler struct {
	Deps *handler.HandlerDeps
}

func NewUpdateFacetHandler(deps *handler.HandlerDeps) *UpdateFacetHandler {
	return &UpdateFacetHandler{deps}
}

// UpdateFacetHandler godoc
// @Summary Update facet
// @Description Updates a facet
// @Tags Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param facetId path uint8 true "Facet ID"
// @Param facet body dto.UpdateFacetInput true "Facet"
// @Success 200 {object} dto.UpdateFacetResponse
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/facets/{facetId} [patch]
func (h *UpdateFacetHandler) Handle(c *gin.Context) {
	var input dto.UpdateFacetInput

	if err := c.ShouldBindUri(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	userCommunicationServiceRepo := repository.NewUserCommunicationServiceRepository(h.Deps.DB)
	userCommunicationServices, err := userCommunicationServiceRepo.FindManyByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to find user communication services")
		return
	}

	validServiceIDs := make(map[uint8]bool)
	for _, ucs := range userCommunicationServices {
		validServiceIDs[ucs.ID] = true
	}

	for _, item := range input.Configuration.Items {
		if !validServiceIDs[uint8(item.Id)] {
			handler.NotFound(c, fmt.Sprintf("Communication service with ID %d not found", item.Id))
			return
		}
	}

	facetRepo := repository.NewFacetRepository(h.Deps.DB)

	facet, err := facetRepo.FindOneByIDAndUserID(input.FacetID, userId)
	if err != nil {
		handler.NotFound(c, "Facet not found")
		return
	}

	if input.Color != nil {
		facet.Color = *input.Color
	}
	if input.PublicLabel != nil {
		facet.PublicLabel = *input.PublicLabel
	}
	if input.PrivateLabel != nil {
		facet.PrivateLabel = *input.PrivateLabel
	}
	if input.Configuration != nil {
		facet.Configuration = *input.Configuration
	}

	err = facetRepo.UpdateOne(facet)
	if err != nil {
		handler.InternalError(c, "Failed to update facet: "+err.Error())
		return
	}

	handler.OK(c, dto.ToUpdateFacetResponse(*facet, userCommunicationServices))
}
