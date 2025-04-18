package handler

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/request"
	"hs-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// GetUserMeHandler godoc
// @Summary Get user me
// @Description Returns user me
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.UserResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/users/me [get]
func (h *UserHandler) GetUserMe(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	user, err := h.UserService.GetUserById(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, user)
}

// GetUserByEmailHandler godoc
// @Summary Get user by email
// @Description Returns user info by email query param
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param email path string true "Email address"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Router /api/users/email/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	var request request.GetUserByEmailRequest

	if err := c.ShouldBindUri(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	user, err := h.UserService.GetUserByEmail(request.Email)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, user)
}
