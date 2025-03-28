package user

import (
	"net/http"

	dto "hs-backend/internal/dto/user"
	"hs-backend/internal/error"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

type GetUserByEmailHandler struct {
	UserRepository repository.UserRepository
}

func NewGetUserByEmailHandler(userRepository repository.UserRepository) *GetUserByEmailHandler {
	return &GetUserByEmailHandler{userRepository}
}

// GetUserByEmailHandler godoc
// @Summary Get user by email
// @Description Returns user info by email query param
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param email query string true "Email address"
// @Success 200 {object} dto.GetUserByEmailResponse
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Router /api/users [get]
func (h *GetUserByEmailHandler) Handle(c *gin.Context) {
	var input dto.GetUserByEmailInput

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, error.ErrorResponse{Error: "Invalid email format"})
		return
	}

	userRepository := h.UserRepository

	u, err := userRepository.FindOneByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, error.ErrorResponse{Error: "User not found"})
		return
	}

	userDTO := &dto.GetUserByEmailResponse{
		ID:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		AvatarURL: func() string {
			if u.AvatarURL != nil {
				return *u.AvatarURL
			}
			return ""
		}(),
	}

	c.JSON(http.StatusOK, userDTO)
}
