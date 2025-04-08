package prism

import (
	"errors"
	"net/http"

	dto "hs-backend/internal/dto/prism"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeletePrismHandler struct {
	PrismRepository repository.PrismRepository
}

func NewDeletePrismHandler(prismRepository repository.PrismRepository) *DeletePrismHandler {
	return &DeletePrismHandler{prismRepository}
}

// DeletePrismHandler godoc
// @Summary Delete prism
// @Description Deletes a prism
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prismId path uint8 true "Prism ID"
// @Success 204
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/prisms/{prismId} [delete]
func (h *DeletePrismHandler) Handle(c *gin.Context) {
	prismRepository := h.PrismRepository

	var input dto.DeletePrismInput

	if err := c.ShouldBindUri(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))
	if err := prismRepository.DeleteOneByIDAndUserID(input.PrismID, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.NotFound(c, "Prism not found")
			return
		}
		handler.InternalError(c, "Failed to delete prism: "+err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
