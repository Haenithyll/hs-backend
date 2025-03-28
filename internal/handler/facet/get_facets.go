package facet

import (
	dto "hs-backend/internal/dto/facet"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetFacetsHandler struct {
	FacetRepository                    repository.FacetRepository
	UserCommunicationServiceRepository repository.UserCommunicationServiceRepository
}

func NewGetFacetsHandler(facetRepository repository.FacetRepository, userCommunicationServiceRepository repository.UserCommunicationServiceRepository) *GetFacetsHandler {
	return &GetFacetsHandler{facetRepository, userCommunicationServiceRepository}
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
	facetRepository := h.FacetRepository
	userCommunicationServiceRepository := h.UserCommunicationServiceRepository

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	facets, err := facetRepository.FindManyByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to get facets: "+err.Error())
		return
	}

	userCommunicationServices, err := userCommunicationServiceRepository.FindManyByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to find user communication services")
		return
	}
	handler.OK(c, dto.ToGetFacetsResponse(facets, userCommunicationServices))
}
