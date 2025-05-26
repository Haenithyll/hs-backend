package handler

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/request"
	"hs-backend/internal/request/validator"
	"hs-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestHandler struct {
	RequestService *service.RequestService
}

func NewRequestHandler(requestService *service.RequestService) *RequestHandler {
	return &RequestHandler{RequestService: requestService}
}

// GetAllReceivedRequestsHandler godoc
// @Summary Get received requests
// @Description Returns received requests
// @Tags Requests
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.RequestResponses
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/requests/received [get]
func (h *RequestHandler) GetAllReceived(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	requests, err := h.RequestService.GetReceivedRequests(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, requests)
}

// GetAllIssuedRequestsHandler godoc
// @Summary Get issued requests
// @Description Returns issued requests
// @Tags Requests
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.RequestResponses
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/requests/issued [get]
func (h *RequestHandler) GetAllIssued(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	requests, err := h.RequestService.GetIssuedRequests(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, requests)
}

// CreateRequestHandler godoc
// @Summary Create request
// @Description Creates a new request
// @Tags Requests
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.CreateRequestRequest true "Request"
// @Success 200 {object} response.RequestResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/requests [post]
func (h *RequestHandler) Create(c *gin.Context) {
	var input request.CreateRequestRequest

	if err := c.ShouldBind(&input); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := validator.ValidateCreateRequestRequest(&input); err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	request, err := h.RequestService.CreateRequest(userId, input)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, request)
}

// MarkRequestAsReadHandler godoc
// @Summary Mark request as read
// @Description Marks a request as read
// @Tags Requests
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param requestId path uint8 true "Request ID"
// @Success 204
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/requests/{requestId}/read [put]
func (h *RequestHandler) MarkAsRead(c *gin.Context) {
	var input request.MarkRequestAsReadRequest

	if err := c.ShouldBindUri(&input); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	err := h.RequestService.MarkRequestAsRead(userId, input)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.NoContent(c)
}
