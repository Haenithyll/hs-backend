package user

import (
	dto "hs-backend/internal/dto/user"
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetUserMeHandler struct {
	UserRepository repository.UserRepository
}

func NewGetUserMeHandler(userRepository repository.UserRepository) *GetUserMeHandler {
	return &GetUserMeHandler{userRepository}
}

// GetUserMeHandler godoc
// @Summary Get user me
// @Description Returns user me
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetUserMeResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/users/me [get]
func (h *GetUserMeHandler) Handle(c *gin.Context) {
	userRepository := h.UserRepository

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	user, err := userRepository.FindOneById(userId)
	if err != nil {
		handler.InternalError(c, "Failed to get user: "+err.Error())
		return
	}

	handler.OK(c, dto.ToGetUserMeResponse(user))
}
