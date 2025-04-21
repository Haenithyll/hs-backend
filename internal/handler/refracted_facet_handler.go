package handler

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RefractedFacetHandler struct {
	RefractedFacetService service.RefractedFacetService
}

func NewRefractedFacetHandler(refractedFacetService service.RefractedFacetService) *RefractedFacetHandler {
	return &RefractedFacetHandler{refractedFacetService}
}

// GetRefractedFacetsHandler godoc
// @Summary Get refracted facets
// @Description Returns refracted facets
// @Tags Refracted Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.RefractedFacetResponses
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/refracted-facets [get]
func (h *RefractedFacetHandler) GetAll(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	refractedFacets, err := h.RefractedFacetService.GetRefractedFacets(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, refractedFacets)
}
