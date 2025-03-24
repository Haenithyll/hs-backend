package user_communication_service

import (
	dto "hs-backend/internal/dto/user_communication_service"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUserCommunicationServiceHandler struct {
	Deps *handler.HandlerDeps
}

func NewUpdateUserCommunicationServiceHandler(deps *handler.HandlerDeps) *UpdateUserCommunicationServiceHandler {
	return &UpdateUserCommunicationServiceHandler{deps}
}

// UpdateUserCommunicationServiceHandler godoc
// @Summary Update user communication service
// @Description Updates a user communication service
// @Tags User Communication Services
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userCommunicationServiceId path uint8 true "User Communication Service ID"
// @Param userCommunicationService body dto.UpdateUserCommunicationServiceInput true "User Communication Service"
// @Success 200 {object} dto.UpdateUserCommunicationServiceResponse
// @Failure 400 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/users/communication-services/{userCommunicationServiceId} [patch]
func (h *UpdateUserCommunicationServiceHandler) Handle(c *gin.Context) {
	var input dto.UpdateUserCommunicationServiceInput

	if err := c.ShouldBindUri(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	if err := c.ShouldBind(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	userCommunicationServiceRepo := repository.NewUserCommunicationServiceRepository(h.Deps.DB)

	ucs, err := userCommunicationServiceRepo.FindByIDAndUserID(input.UserCommunicationServiceID, userId)
	if err != nil {
		handler.NotFound(c, "User communication service not found")
		return
	}

	if input.Name != nil {
		ucs.Name = *input.Name
	}
	if input.Value != nil {
		ucs.Value = *input.Value
	}
	if input.ServiceType != nil {
		ucs.Service = *input.ServiceType
	}

	err = userCommunicationServiceRepo.Update(ucs)
	if err != nil {
		handler.InternalError(c, "Failed to update user communication service: "+err.Error())
		return
	}

	handler.OK(c, dto.ToUpdateUserCommunicationServiceResponse(*ucs))
}
