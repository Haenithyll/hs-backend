package handler

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/request"
	"hs-backend/internal/request/validator"
	"hs-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserCommunicationServiceHandler struct {
	UserCommunicationServiceService service.UserCommunicationServiceService
}

func NewUserCommunicationServiceHandler(
	userCommunicationServiceService *service.UserCommunicationServiceService,
) *UserCommunicationServiceHandler {
	return &UserCommunicationServiceHandler{
		UserCommunicationServiceService: *userCommunicationServiceService,
	}
}

// GetUserCommunicationServiceHandler godoc
// @Summary Get user communication services
// @Description Returns user communication services
// @Tags User Communication Services
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.UserCommunicationServiceResponses
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/users/communication-services [get]
func (h *UserCommunicationServiceHandler) GetAll(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	userCommunicationServices, err := h.UserCommunicationServiceService.GetAllUserCommunicationServices(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, userCommunicationServices)
}

// CreateUserCommunicationServiceHandler godoc
// @Summary Create user communication service
// @Description Creates a new user communication service
// @Tags User Communication Services
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userCommunicationService body request.CreateUserCommunicationServiceRequest true "User Communication Service"
// @Success 200 {object} response.UserCommunicationServiceResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/users/communication-services [post]
func (h *UserCommunicationServiceHandler) Create(c *gin.Context) {
	var input request.CreateUserCommunicationServiceRequest

	if err := c.ShouldBind(&input); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := validator.ValidateCreateUserCommunicationServiceRequest(&input); err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	userCommunicationService, err := h.UserCommunicationServiceService.CreateUserCommunicationService(userId, input)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, userCommunicationService)
}

// UpdateUserCommunicationServiceHandler godoc
// @Summary Update user communication service
// @Description Updates a user communication service
// @Tags User Communication Services
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userCommunicationServiceId path uint8 true "User Communication Service ID"
// @Param userCommunicationService body request.UpdateUserCommunicationServiceRequest true "User Communication Service"
// @Success 200 {object} response.UserCommunicationServiceResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/users/communication-services/{userCommunicationServiceId} [patch]
func (h *UserCommunicationServiceHandler) Update(c *gin.Context) {
	var request request.UpdateUserCommunicationServiceRequest

	if err := c.ShouldBindUri(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := validator.ValidateUpdateUserCommunicationServiceRequest(&request); err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	userCommunicationService, err := h.UserCommunicationServiceService.UpdateUserCommunicationService(userId, request)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, userCommunicationService)
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
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/users/communication-services/{userCommunicationServiceId} [delete]
func (h *UserCommunicationServiceHandler) Delete(c *gin.Context) {
	var request request.DeleteUserCommunicationServiceRequest

	if err := c.ShouldBindUri(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	err := h.UserCommunicationServiceService.DeleteUserCommunicationService(userId, request.UserCommunicationServiceID)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.NoContent(c)
}
