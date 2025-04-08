package prism

import (
	"fmt"
	"net/http"

	dto "hs-backend/internal/dto/prism"
	"hs-backend/internal/handler"
	"hs-backend/internal/model"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdatePrismHandler struct {
	FacetRepository repository.FacetRepository
	PrismRepository repository.PrismRepository
}

func NewUpdatePrismHandler(
	facetRepository repository.FacetRepository,
	prismRepository repository.PrismRepository,
) *UpdatePrismHandler {
	return &UpdatePrismHandler{facetRepository, prismRepository}
}

// UpdatePrismHandler godoc
// @Summary Update prism
// @Description Updates a prism
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prismId path uint8 true "Prism ID"
// @Param prism body dto.UpdatePrismInput true "Prism"
// @Success 200 {object} dto.UpdatePrismResponse
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/prisms/{prismId} [patch]
func (h *UpdatePrismHandler) Handle(c *gin.Context) {
	facetRepository := h.FacetRepository
	prismRepository := h.PrismRepository

	var input dto.UpdatePrismInput

	if err := c.ShouldBindUri(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	prism, err := prismRepository.FindOneByIDAndUserID(input.PrismID, userId)
	if err != nil {
		handler.NotFound(c, "Prism not found")
		return
	}

	facets, err := facetRepository.FindManyByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to find facets")
		return
	}

	if input.Configuration != nil {
		if err := h.validateConfiguration(input, facets); err != nil {
			handler.BadRequest(c, err.Error())
			return
		}
		prism.Configuration = *input.Configuration
	}

	if input.Name != nil {
		prism.Name = *input.Name
	}

	err = prismRepository.UpdateOne(prism)
	if err != nil {
		handler.InternalError(c, "Failed to update prism: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ToUpdatePrismResponse(*prism, facets))
}

func (h *UpdatePrismHandler) validateConfiguration(input dto.UpdatePrismInput, facets []model.Facet) error {
	validFacetIDs := make(map[uint8]bool)

	for _, facet := range facets {
		validFacetIDs[facet.ID] = true
	}

	if !validFacetIDs[input.Configuration.Base] {
		return fmt.Errorf("facet with ID %d not found", input.Configuration.Base)
	}

	for _, item := range input.Configuration.Users {
		if !validFacetIDs[item.FacetId] {
			return fmt.Errorf("facet with ID %d not found", item.FacetId)
		}
	}

	return nil
}
