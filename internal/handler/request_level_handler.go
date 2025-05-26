package handler

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type RequestLevelHandler struct {
	RequestLevelService *service.RequestLevelService
}

func NewRequestLevelHandler(requestLevelService *service.RequestLevelService) *RequestLevelHandler {
	return &RequestLevelHandler{RequestLevelService: requestLevelService}
}

// GetAllRequestLevelsHandler godoc
// @Summary Get all request levels
// @Description Returns all request levels
// @Tags Request Levels
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.RequestLevelResponses
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/requests/levels [get]
func (h *RequestLevelHandler) GetAll(c *gin.Context) {
	requestLevels, err := h.RequestLevelService.GetAllRequestLevels()
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, requestLevels)
}
