package user_communication_service

import (
	"net/http"

	dto "hs-backend/internal/dto/user_communication_service"
	"hs-backend/internal/handler"
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserCommunicationServiceHandler struct {
	UserCommunicationServiceRepository repository.UserCommunicationServiceRepository
}

func NewCreateUserCommunicationServiceHandler(
	userCommunicationServiceRepository repository.UserCommunicationServiceRepository,
) *CreateUserCommunicationServiceHandler {
	return &CreateUserCommunicationServiceHandler{userCommunicationServiceRepository}
}

// CreateUserCommunicationServiceHandler godoc
// @Summary Create user communication service
// @Description Creates a new user communication service
// @Tags User Communication Services
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userCommunicationService body dto.CreateUserCommunicationServiceInput true "User Communication Service"
// @Success 200 {object} dto.CreateUserCommunicationServiceResponse
// @Failure 400 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/users/communication-services [post]
func (h *CreateUserCommunicationServiceHandler) Handle(c *gin.Context) {
	userCommunicationServiceRepository := h.UserCommunicationServiceRepository

	var input dto.CreateUserCommunicationServiceInput

	if err := c.ShouldBind(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	ucs := model.UserCommunicationService{
		UserId:  userId,
		Name:    input.Name,
		Value:   input.Value,
		Service: enum.CommunicationService(input.Service),
	}

	err := userCommunicationServiceRepository.CreateOne(&ucs)
	if err != nil {
		handler.InternalError(c, "Failed to create user communication service: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ToCreateUserCommunicationServiceResponse(ucs))
}
