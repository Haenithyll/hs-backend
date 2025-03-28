package facet

import (
	"errors"
	"net/http"

	dto "hs-backend/internal/dto/facet"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeleteFacetHandler struct {
	Deps *handler.HandlerDeps
}

func NewDeleteFacetHandler(deps *handler.HandlerDeps) *DeleteFacetHandler {
	return &DeleteFacetHandler{deps}
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
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/facets/{facetId} [delete]
func (h *DeleteFacetHandler) Handle(c *gin.Context) {
	var input dto.DeleteFacetInput

	if err := c.ShouldBindUri(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	repo := repository.NewFacetRepository(h.Deps.DB)

	userId := uuid.MustParse(c.MustGet("user_id").(string))
	if err := repo.DeleteOneByIDAndUserID(input.FacetID, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.NotFound(c, "Facet not found")
			return
		}
		handler.InternalError(c, "Failed to delete facet: "+err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
