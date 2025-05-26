package handler

import (
	"hs-backend/internal/domain"
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

// GetAllUsersHandler godoc
// @Summary Get all users
// @Description Returns all users
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.UserResponses
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, users)
}

// GetMeUserHandler godoc
// @Summary Get my user information
// @Description Returns my user information
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.UserResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/users/me [get]
func (h *UserHandler) GetMe(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	user, err := h.UserService.GetUserById(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, user)
}
