package facet

import (
	"fmt"
	"net/http"

	dto "hs-backend/internal/dto/facet"
	"hs-backend/internal/handler"
	"hs-backend/internal/model"
	"hs-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateFacetHandler struct {
	Deps *handler.HandlerDeps
}

func NewCreateFacetHandler(deps *handler.HandlerDeps) *CreateFacetHandler {
	return &CreateFacetHandler{deps}
}

// CreateFacetHandler godoc
// @Summary Create facet
// @Description Creates a new facet
// @Tags Facets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param facet body dto.CreateFacetInput true "Facet"
// @Success 200 {object} dto.CreateFacetResponse
// @Failure 400 {object} error.ErrorResponse
// @Failure 404 {object} error.ErrorResponse
// @Failure 500 {object} error.ErrorResponse
// @Router /api/facets [post]
func (h *CreateFacetHandler) Handle(c *gin.Context) {
	var input dto.CreateFacetInput

	if err := c.ShouldBind(&input); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		handler.BadRequest(c, err.Error())
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	userCommunicationServiceRepo := repository.NewUserCommunicationServiceRepository(h.Deps.DB)
	userCommunicationServiceIds, err := userCommunicationServiceRepo.FindIDsByUserId(userId)
	if err != nil {
		handler.InternalError(c, "Failed to find user communication services")
		return
	}

	validServiceIDs := make(map[uint8]bool)
	for _, id := range userCommunicationServiceIds {
		validServiceIDs[id] = true
	}

	for _, item := range input.Configuration.Items {
		if !validServiceIDs[uint8(item.Id)] {
			handler.NotFound(c, fmt.Sprintf("Communication service with ID %d not found", item.Id))
			return
		}
	}

	repo := repository.NewFacetRepository(h.Deps.DB)

	facet := model.Facet{
		Color:         input.Color,
		PublicLabel:   input.PublicLabel,
		PrivateLabel:  input.PrivateLabel,
		Configuration: input.Configuration,
		UserId:        userId,
	}

	err = repo.Create(&facet)
	if err != nil {
		handler.InternalError(c, "Failed to create facet: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.ToCreateFacetResponse(facet))
}
