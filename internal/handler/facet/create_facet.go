package facet

import (
	"fmt"
	"net/http"

	dto "hs-backend/internal/dto/facet"
	"hs-backend/internal/error"
	"hs-backend/internal/handler"
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
// @Failure 500 {object} error.ErrorResponse
// @Router /api/facets [post]
func (h *CreateFacetHandler) Handle(c *gin.Context) {
	var input dto.CreateFacetInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, error.ErrorResponse{Error: "Invalid input format"})
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, error.ErrorResponse{Error: err.Error()})
		return
	}

	userId := uuid.MustParse(c.MustGet("user_id").(string))

	userCommunicationServiceRepo := repository.NewUserCommunicationServiceRepository(h.Deps.DB)
	userCommunicationServiceIds, err := userCommunicationServiceRepo.FindIDsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, error.ErrorResponse{Error: "Failed to find user communication services"})
		return
	}

	validServiceIDs := make(map[uint8]bool)
	for _, id := range userCommunicationServiceIds {
		validServiceIDs[id] = true
	}

	for _, item := range input.Configuration.Items {
		if !validServiceIDs[uint8(item.Id)] {
			c.JSON(http.StatusBadRequest, error.ErrorResponse{
				Error: fmt.Sprintf("Invalid communication service ID: %d", item.Id),
			})
			return
		}
	}

	repo := repository.NewFacetRepository(h.Deps.DB)

	f, err := repo.Create(input, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, error.ErrorResponse{Error: "Failed to create facet"})
		return
	}

	facetDTO := &dto.CreateFacetResponse{
		Facet: dto.ToFacetResponse(*f),
	}

	c.JSON(http.StatusOK, facetDTO)
}
