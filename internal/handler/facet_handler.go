package handler

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/request"
	"hs-backend/internal/request/validator"
	"hs-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FacetHandler struct {
	FacetService *service.FacetService
}

func NewFacetHandler(facetService *service.FacetService) *FacetHandler {
	return &FacetHandler{FacetService: facetService}
}

// GetFacetsHandler godoc
// @Summary Get facets
// @Description Returns facets
// @Tags Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.FacetResponses
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/facets [get]
func (h *FacetHandler) GetAll(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	facets, err := h.FacetService.GetFacets(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, facets)
}

// CreateFacetHandler godoc
// @Summary Create facet
// @Description Creates a new facet
// @Tags Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param facet body request.CreateFacetRequest true "Facet"
// @Success 200 {object} response.FacetResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/facets [post]
func (h *FacetHandler) Create(c *gin.Context) {
	var request request.CreateFacetRequest

	if err := c.ShouldBind(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := validator.ValidateCreateFacetRequest(request); err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	facet, err := h.FacetService.CreateFacet(userId, request)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, facet)
}

// UpdateFacetHandler godoc
// @Summary Update facet
// @Description Updates a facet
// @Tags Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param facetId path uint8 true "Facet ID"
// @Param facet body request.UpdateFacetRequest true "Facet"
// @Success 200 {object} response.FacetResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/facets/{facetId} [patch]
func (h *FacetHandler) Update(c *gin.Context) {
	var request request.UpdateFacetRequest

	if err := c.ShouldBindUri(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := validator.ValidateUpdateFacetRequest(request); err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	facet, err := h.FacetService.UpdateFacet(userId, request)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, facet)
}

// DeleteFacetHandler godoc
// @Summary Delete facet
// @Description Deletes a facet
// @Tags Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param facetId path uint8 true "Facet ID"
// @Success 204
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/facets/{facetId} [delete]
func (h *FacetHandler) Delete(c *gin.Context) {
	var request request.DeleteFacetRequest

	if err := c.ShouldBindUri(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	err := h.FacetService.DeleteFacet(userId, request.FacetID)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.NoContent(c)
}
