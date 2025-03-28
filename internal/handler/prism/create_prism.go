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

type CreatePrismHandler struct {
	Deps *handler.HandlerDeps
}

func NewCreatePrismHandler(deps *handler.HandlerDeps) *CreatePrismHandler {
	return &CreatePrismHandler{deps}
}

// CreatePrismHandler godoc
// @Summary Create prism
// @Description Creates a new prism
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prism body dto.CreatePrismInput true "Prism"
// @Success 200 {object} dto.CreatePrismResponse
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/prisms [post]
func (h *CreatePrismHandler) Handle(c *gin.Context) {
	var input dto.CreatePrismInput

	if err := c.ShouldBind(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	facetRepo := repository.NewFacetRepository(h.Deps.DB)
	facets, err := facetRepo.FindManyByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to find facets")
		return
	}

	validFacetIDs := make(map[uint8]bool)
	for _, facet := range facets {
		validFacetIDs[facet.ID] = true
	}

	if !validFacetIDs[input.Configuration.Base] {
		handler.NotFound(c, fmt.Sprintf("Facet with ID %d not found", input.Configuration.Base))
		return
	}

	for _, item := range input.Configuration.Users {
		if !validFacetIDs[item.FacetId] {
			handler.NotFound(c, fmt.Sprintf("Facet with ID %d not found", item.FacetId))
			return
		}
	}

	prismRepo := repository.NewPrismRepository(h.Deps.DB)

	prism := model.Prism{
		Name:          input.Name,
		Configuration: input.Configuration,
		UserId:        userId,
	}

	err = prismRepo.CreateOne(&prism)
	if err != nil {
		handler.InternalError(c, "Failed to create prism: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ToCreatePrismResponse(prism, facets))
}
