package facet

import (
	dto "hs-backend/internal/dto/facet"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetFacetsHandler struct {
	Deps *handler.HandlerDeps
}

func NewGetFacetsHandler(deps *handler.HandlerDeps) *GetFacetsHandler {
	return &GetFacetsHandler{deps}
}

// GetFacetsHandler godoc
// @Summary Get facets
// @Description Returns facets
// @Tags Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetFacetsResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/facets [get]
func (h *GetFacetsHandler) Handle(c *gin.Context) {
	repo := repository.NewFacetRepository(h.Deps.DB)

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	facets, err := repo.FindByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to get facets: "+err.Error())
		return
	}

	handler.OK(c, dto.ToGetFacetsResponse(facets))
}
