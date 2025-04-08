package user_communication_service

import (
	dto "hs-backend/internal/dto/user_communication_service"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetUserCommunicationServiceHandler struct {
	UserCommunicationServiceRepository repository.UserCommunicationServiceRepository
}

func NewGetUserCommunicationServiceHandler(
	userCommunicationServiceRepository repository.UserCommunicationServiceRepository,
) *GetUserCommunicationServiceHandler {
	return &GetUserCommunicationServiceHandler{userCommunicationServiceRepository}
}

// GetUserCommunicationServiceHandler godoc
// @Summary Get user communication services
// @Description Returns user communication services
// @Tags User Communication Services
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetUserCommunicationServicesResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/users/communication-services [get]
func (h *GetUserCommunicationServiceHandler) Handle(c *gin.Context) {
	userCommunicationServiceRepository := h.UserCommunicationServiceRepository

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	ucsList, err := userCommunicationServiceRepository.FindManyByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to get user communication services: "+err.Error())
		return
	}

	handler.OK(c, dto.ToGetUserCommunicationServicesResponse(ucsList))
}
