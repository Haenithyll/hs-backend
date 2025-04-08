package prism

import (
	"net/http"

	dto "hs-backend/internal/dto/prism"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ActivatePrismHandler struct {
	PrismRepository            repository.PrismRepository
	UserPrismTrackerRepository repository.UserPrismTrackerRepository
}

func NewActivatePrismHandler(
	prismRepository repository.PrismRepository,
	userPrismTrackerRepository repository.UserPrismTrackerRepository,
) *ActivatePrismHandler {
	return &ActivatePrismHandler{prismRepository, userPrismTrackerRepository}
}

// ActivatePrismHandler godoc
// @Summary Activate prism
// @Description Activates a prism
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prismId path uint8 true "Prism ID"
// @Success 204
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/prisms/{prismId}/activate [post]
func (h *ActivatePrismHandler) Handle(c *gin.Context) {
	prismRepository := h.PrismRepository
	userPrismTrackerRepository := h.UserPrismTrackerRepository

	var input dto.ActivatePrismInput

	if err := c.ShouldBindUri(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	if _, err := prismRepository.FindOneByIDAndUserID(input.PrismID, userId); err != nil {
		handler.NotFound(c, "Prism not found")
		return
	}

	if err := userPrismTrackerRepository.ActivatePrismByPrismIdAndUserId(input.PrismID, userId); err != nil {
		handler.InternalError(c, "Failed to activate prism: "+err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
