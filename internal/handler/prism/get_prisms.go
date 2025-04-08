package prism

import (
	dto "hs-backend/internal/dto/prism"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetPrismsHandler struct {
	FacetRepository repository.FacetRepository
	PrismRepository repository.PrismRepository
}

func NewGetPrismsHandler(
	facetRepository repository.FacetRepository,
	prismRepository repository.PrismRepository,
) *GetPrismsHandler {
	return &GetPrismsHandler{facetRepository, prismRepository}
}

// GetPrismsHandler godoc
// @Summary Get prisms
// @Description Returns prisms
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetPrismsResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/prisms [get]
func (h *GetPrismsHandler) Handle(c *gin.Context) {
	facetRepository := h.FacetRepository
	prismRepository := h.PrismRepository

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	prisms, err := prismRepository.FindManyByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to get prisms: "+err.Error())
		return
	}

	facets, err := facetRepository.FindManyByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to get facets: "+err.Error())
		return
	}

	handler.OK(c, dto.ToGetPrismsResponse(prisms, facets))
}
