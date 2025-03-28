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
	FacetRepository                    repository.FacetRepository
	UserCommunicationServiceRepository repository.UserCommunicationServiceRepository
}

func NewUpdateFacetHandler(facetRepository repository.FacetRepository, userCommunicationServiceRepository repository.UserCommunicationServiceRepository) *UpdateFacetHandler {
	return &UpdateFacetHandler{facetRepository, userCommunicationServiceRepository}
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
	facetRepository := h.FacetRepository
	userCommunicationServiceRepository := h.UserCommunicationServiceRepository

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

	userCommunicationServices, err := userCommunicationServiceRepository.FindManyByUserId(userId)
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

	facet, err := facetRepository.FindOneByIDAndUserID(input.FacetID, userId)
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

	err = facetRepository.UpdateOne(facet)
	if err != nil {
		handler.InternalError(c, "Failed to update facet: "+err.Error())
		return
	}

	handler.OK(c, dto.ToUpdateFacetResponse(*facet, userCommunicationServices))
}
