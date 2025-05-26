package handler

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type KeyHandler struct {
	KeyService *service.KeyService
}

func NewKeyHandler(keyService *service.KeyService) *KeyHandler {
	return &KeyHandler{KeyService: keyService}
}

// GenerateKeyHandler godoc
// @Summary Generate a new key
// @Description Generates a new key
// @Tags Key
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.KeyResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/key/generate [post]
func (h *KeyHandler) Generate(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	key, err := h.KeyService.GenerateKey(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, key)
}
