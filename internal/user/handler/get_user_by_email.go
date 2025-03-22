package handler

import (
	"net/http"

	"hs-backend/internal/error"
	"hs-backend/internal/shared"
	"hs-backend/internal/user"
	"hs-backend/internal/user/dto"
	"hs-backend/internal/user/query"

	"github.com/gin-gonic/gin"
)

type GetUserByEmailHandler struct {
	Deps *shared.HandlerDeps
}

func NewGetUserByEmailHandler(deps *shared.HandlerDeps) *GetUserByEmailHandler {
	return &GetUserByEmailHandler{deps}
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

	repo := user.NewRepository(h.Deps.DB)
	handler := query.NewGetUserByEmailHandler(repo)

	userDTO, err := handler.Handle(input)
	if err != nil {
		c.JSON(http.StatusNotFound, error.ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, userDTO)
}
