package user_communication_service

import (
	"errors"
	"net/http"

	dto "hs-backend/internal/dto/user_communication_service"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeleteUserCommunicationServiceHandler struct {
	Deps *handler.HandlerDeps
}

func NewDeleteUserCommunicationServiceHandler(deps *handler.HandlerDeps) *DeleteUserCommunicationServiceHandler {
	return &DeleteUserCommunicationServiceHandler{deps}
}

// DeleteUserCommunicationServiceHandler godoc
// @Summary Delete user communication service
// @Description Deletes a user communication service
// @Tags User Communication Services
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userCommunicationServiceId path uint8 true "User Communication Service ID"
// @Success 204
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/users/communication-services/{userCommunicationServiceId} [delete]
func (h *DeleteUserCommunicationServiceHandler) Handle(c *gin.Context) {
	var input dto.DeleteUserCommunicationServiceInput

	if err := c.ShouldBindUri(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	repo := repository.NewUserCommunicationServiceRepository(h.Deps.DB)

	userId := uuid.MustParse(c.MustGet("user_id").(string))
	if err := repo.DeleteByIDAndUserID(input.UserCommunicationServiceID, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.NotFound(c, "User communication service not found")
			return
		}
		handler.InternalError(c, "Failed to delete user communication service: "+err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
