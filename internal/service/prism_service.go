package service

import (
	"errors"
	"fmt"
	"hs-backend/internal/domain"
	"hs-backend/internal/model"
	"hs-backend/internal/repository"
	"hs-backend/internal/request"
	"hs-backend/internal/response"
	"hs-backend/internal/response/mapper"
	"hs-backend/internal/util/map_util"
	"hs-backend/internal/validation"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrismService struct {
	PrismRepository            repository.PrismRepository
	FacetRepository            repository.FacetRepository
	UserPrismTrackerRepository repository.UserPrismTrackerRepository

	UserRepository repository.UserRepository
}

func NewPrismService(
	prismRepository repository.PrismRepository,
	facetRepository repository.FacetRepository,
	userPrismTrackerRepository repository.UserPrismTrackerRepository,
	userRepository repository.UserRepository,
) *PrismService {
	return &PrismService{prismRepository, facetRepository, userPrismTrackerRepository, userRepository}
}

// Fetches all prisms associated with the given user ID.
// It also retrieves related facet and user data for DTO conversion.
func (s *PrismService) GetPrisms(userID uuid.UUID) ([]response.PrismResponse, *domain.DomainError) {
	prisms, err := s.PrismRepository.FindManyByUserId(userID)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get prisms: "+err.Error())
	}

	facets, err := s.getFacetsFromPrismConfigs(prisms, userID)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get facets: "+err.Error())
	}

	users, err := s.getUsersFromPrismConfigs(prisms)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get users: "+err.Error())
	}

	facetMap := map_util.BuildFacetMapById(facets)
	userMap := map_util.BuildUserMapById(users)

	return mapper.ToPrismResponses(prisms, facetMap, userMap), nil
}

// Creates a new prism for the given user
// It also checks if the facet IDs in the prism configuration are valid
// and retrieves the corresponding facet and user models for DTO conversion.
func (s *PrismService) CreatePrism(userId uuid.UUID, request request.CreatePrismRequest) (*response.PrismResponse, *domain.DomainError) {
	facets, err := s.FacetRepository.FindManyByUserId(userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "Failed to find facets")
	}

	if err := validation.ValidateConfigurationFacetIds(request.Configuration, facets); err != nil {
		return nil, err
	}

	newPrism := model.Prism{
		Name:          request.Name,
		Configuration: request.Configuration,
		UserId:        userId,
	}

	err = s.PrismRepository.CreateOne(&newPrism)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "Failed to create prism: "+err.Error())
	}

	users, err := s.getUsersFromPrismConfigs([]model.Prism{newPrism})
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get users: "+err.Error())
	}

	facetMap := map_util.BuildFacetMapById(facets)
	userMap := map_util.BuildUserMapById(users)

	response := mapper.ToPrismResponse(newPrism, facetMap, userMap)

	return &response, nil
}

// Updates a prism for the given user
// It also checks if the facet IDs in the prism configuration are valid
// and retrieves the corresponding facet and user models for DTO conversion.
func (s *PrismService) UpdatePrism(userId uuid.UUID, request request.UpdatePrismRequest) (*response.PrismResponse, *domain.DomainError) {
	prism, err := s.PrismRepository.FindOneByIDAndUserID(request.PrismID, userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "Prism not found")
	}

	facets, err := s.FacetRepository.FindManyByUserId(userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "Failed to find facets")
	}

	if request.Configuration != nil {
		if err := validation.ValidateConfigurationFacetIds(*request.Configuration, facets); err != nil {
			return nil, err
		}
		prism.Configuration = *request.Configuration
	}

	if request.Name != nil {
		prism.Name = *request.Name
	}

	err = s.PrismRepository.UpdateOne(prism)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "Failed to update prism: "+err.Error())
	}

	users, err := s.getUsersFromPrismConfigs([]model.Prism{*prism})
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get users: "+err.Error())
	}

	facetMap := map_util.BuildFacetMapById(facets)
	userMap := map_util.BuildUserMapById(users)

	response := mapper.ToPrismResponse(*prism, facetMap, userMap)

	return &response, nil
}

// Deletes a prism for the given user
func (s *PrismService) DeletePrism(userId uuid.UUID, prismId uint8) *domain.DomainError {
	if err := s.PrismRepository.DeleteOneByIDAndUserID(prismId, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.NewDomainError(domain.ErrNotFound, "prism not found")
		}
		return domain.NewDomainError(domain.ErrInternalServerError, "failed to delete prism: "+err.Error())
	}
	return nil
}

// Activates a prism for the given user
func (s *PrismService) ActivatePrism(userId uuid.UUID, prismId uint8) *domain.DomainError {
	if _, err := s.PrismRepository.FindOneByIDAndUserID(prismId, userId); err != nil {
		return domain.NewDomainError(domain.ErrNotFound, "prism not found")
	}

	if err := s.UserPrismTrackerRepository.ActivatePrismByPrismIdAndUserId(prismId, userId); err != nil {
		return domain.NewDomainError(domain.ErrInternalServerError, "failed to activate prism: "+err.Error())
	}

	return nil
}

// Extracts all unique facet IDs from the prism configurations
// and fetches their corresponding facet models for the given user.
func (s *PrismService) getFacetsFromPrismConfigs(prisms []model.Prism, userId uuid.UUID) ([]model.Facet, error) {
	facetIds := make(map[uint8]bool)

	// Map the facet IDs from the prism configurations to ensure uniqueness
	for _, prism := range prisms {
		facetIds[prism.Configuration.Base] = true

		for _, user := range prism.Configuration.Users {
			facetIds[user.FacetId] = true
		}
	}

	// Convert the map to a list of unique facet IDs
	facetIdsList := make([]uint8, 0, len(facetIds))
	for facetId := range facetIds {
		facetIdsList = append(facetIdsList, facetId)
	}

	// Fetch the facets from the repository
	facets, err := s.FacetRepository.FindManyByIdsAndUserId(facetIdsList, userId)
	if err != nil {
		return nil, err
	}

	if len(facets) != len(facetIdsList) {
		return nil, fmt.Errorf("some facets were not found")
	}

	return facets, nil
}

// Extracts all unique user IDs from the prism configurations
// and fetches their corresponding user models.
func (s *PrismService) getUsersFromPrismConfigs(prisms []model.Prism) ([]model.User, error) {
	userIds := make(map[uuid.UUID]bool)

	// Map the user IDs from the prism configurations to ensure uniqueness
	for _, prism := range prisms {
		for _, user := range prism.Configuration.Users {
			userIds[user.UserId] = true
		}
	}

	// Convert the map to a list of unique user IDs
	userIdsList := make([]uuid.UUID, 0, len(userIds))
	for userId := range userIds {
		userIdsList = append(userIdsList, userId)
	}

	// Fetch the users from the repository
	users, err := s.UserRepository.FindManyByIds(userIdsList)
	if err != nil {
		return nil, err
	}

	if len(users) != len(userIdsList) {
		return nil, fmt.Errorf("some users were not found")
	}

	return users, nil
}
