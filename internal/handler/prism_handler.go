package handler

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/request"
	"hs-backend/internal/request/validator"
	"hs-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PrismHandler struct {
	PrismService *service.PrismService
}

func NewPrismHandler(prismService *service.PrismService) *PrismHandler {
	return &PrismHandler{
		PrismService: prismService,
	}
}

// GetPrismsHandler godoc
// @Summary Get prisms
// @Description Returns prisms
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.PrismResponses
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/prisms [get]
func (h *PrismHandler) GetAll(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	prismDtos, err := h.PrismService.GetPrisms(userId)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, prismDtos)
}

// @Summary Create prism
// @Description Creates a new prism
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prism body request.CreatePrismRequest true "Prism"
// @Success 200 {object} response.PrismResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/prisms [post]
func (h *PrismHandler) Create(c *gin.Context) {
	userId := uuid.MustParse(c.MustGet("user_id").(string))

	var input request.CreatePrismRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := validator.ValidateCreatePrismRequest(&input); err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	prism, err := h.PrismService.CreatePrism(userId, input)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, prism)
}

// UpdatePrismHandler godoc
// @Summary Update prism
// @Description Updates a prism
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prismId path uint8 true "Prism ID"
// @Param prism body request.UpdatePrismRequest true "Prism"
// @Success 200 {object} response.PrismResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/prisms/{prismId} [patch]
func (h *PrismHandler) Update(c *gin.Context) {
	var request request.UpdatePrismRequest

	if err := c.ShouldBindUri(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	if err := validator.ValidateUpdatePrismRequest(&request); err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	prism, err := h.PrismService.UpdatePrism(userId, request)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.Ok(c, prism)
}

// DeletePrismHandler godoc
// @Summary Delete prism
// @Description Deletes a prism
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prismId path uint8 true "Prism ID"
// @Success 204
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/prisms/{prismId} [delete]
func (h *PrismHandler) Delete(c *gin.Context) {
	var request request.DeletePrismRequest

	if err := c.ShouldBindUri(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	err := h.PrismService.DeletePrism(userId, request.PrismID)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// ActivatePrismHandler godoc
// @Summary Activate prism
// @Description Activates a prism
// @Tags Prisms
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prismId path uint8 true "Prism ID"
// @Success 204
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /api/prisms/{prismId}/activate [post]
func (h *PrismHandler) Activate(c *gin.Context) {
	var request request.ActivatePrismRequest

	if err := c.ShouldBindUri(&request); err != nil {
		domain.ToErrorResponse(c, domain.NewDomainError(domain.ErrBadRequest, err.Error()))
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	err := h.PrismService.ActivatePrism(userId, request.PrismID)
	if err != nil {
		domain.ToErrorResponse(c, err)
		return
	}

	domain.NoContent(c)
}
