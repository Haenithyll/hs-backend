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

type FacetService struct {
	FacetRepository                    repository.FacetRepository
	UserCommunicationServiceRepository repository.UserCommunicationServiceRepository
}

func NewFacetService(
	facetRepository repository.FacetRepository,
	userCommunicationServiceRepository repository.UserCommunicationServiceRepository,
) *FacetService {
	return &FacetService{facetRepository, userCommunicationServiceRepository}
}

// Fetches all facets associated with the given user ID.
// It also retrieves related user communication services for DTO conversion.
func (s *FacetService) GetFacets(userID uuid.UUID) ([]response.FacetResponse, *domain.DomainError) {
	facets, err := s.FacetRepository.FindManyByUserId(userID)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get facets: "+err.Error())
	}

	userCommunicationServices, err := s.getUserCommunicationServicesFromFacetConfigs(facets, userID)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get user communication services: "+err.Error())
	}

	userCommunicationServiceMap := map_util.BuildUserCommunicationServiceMapById(userCommunicationServices)

	return mapper.ToFacetResponses(facets, userCommunicationServiceMap), nil
}

// Creates a new facet for the given user
// It also checks if the user communication services are valid
// and retrieves the corresponding user communication services for DTO conversion.
func (s *FacetService) CreateFacet(userId uuid.UUID, request request.CreateFacetRequest) (*response.FacetResponse, *domain.DomainError) {
	userCommunicationServices, err := s.UserCommunicationServiceRepository.FindManyByUserId(userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, err.Error())
	}

	if err := validation.ValidateUserCommunicationServiceIds(request.Configuration, userCommunicationServices); err != nil {
		return nil, err
	}

	facet := model.Facet{
		Color:         request.Color,
		PublicLabel:   request.PublicLabel,
		PrivateLabel:  request.PrivateLabel,
		Configuration: request.Configuration,
		UserId:        userId,
	}

	err = s.FacetRepository.CreateOne(&facet)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, err.Error())
	}

	userCommunicationServiceMap := map_util.BuildUserCommunicationServiceMapById(userCommunicationServices)

	reponse := mapper.ToFacetResponse(facet, userCommunicationServiceMap)

	return &reponse, nil
}

// Updates a facet for the given user
// It also checks if the user communication services are valid
// and retrieves the corresponding user communication services for DTO conversion.
func (s *FacetService) UpdateFacet(userId uuid.UUID, request request.UpdateFacetRequest) (*response.FacetResponse, *domain.DomainError) {
	facet, err := s.FacetRepository.FindOneByIDAndUserID(request.FacetID, userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "Facet not found")
	}

	userCommunicationServices, err := s.getUserCommunicationServicesFromFacetConfigs([]model.Facet{*facet}, userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "Failed to find facets")
	}

	if request.Configuration != nil {
		if err := validation.ValidateUserCommunicationServiceIds(*request.Configuration, userCommunicationServices); err != nil {
			return nil, err
		}
		facet.Configuration = *request.Configuration
	}

	err = s.FacetRepository.UpdateOne(facet)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "Failed to update facet: "+err.Error())
	}

	userCommunicationServiceMap := map_util.BuildUserCommunicationServiceMapById(userCommunicationServices)

	response := mapper.ToFacetResponse(*facet, userCommunicationServiceMap)

	return &response, nil
}

// Deletes a facet for the given user
func (s *FacetService) DeleteFacet(userId uuid.UUID, facetId uint8) *domain.DomainError {
	if err := s.FacetRepository.DeleteOneByIDAndUserID(facetId, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.NewDomainError(domain.ErrNotFound, "facet not found")
		}
		return domain.NewDomainError(domain.ErrInternalServerError, "failed to delete facet: "+err.Error())
	}
	return nil
}

// Extracts all unique user communication service IDs from the facet configurations
// and fetches their corresponding user communication service models for the given user.
func (s *FacetService) getUserCommunicationServicesFromFacetConfigs(facets []model.Facet, userId uuid.UUID) ([]model.UserCommunicationService, error) {
	userCommunicationServiceIds := make(map[uint8]bool)

	// Map the user communication service IDs from the facet configurations to ensure uniqueness
	for _, facet := range facets {
		for _, item := range facet.Configuration.Items {
			userCommunicationServiceIds[item.Id] = true
		}
	}

	// Convert the map to a list of unique user communication service IDs
	userCommunicationServiceIdsList := make([]uint8, 0, len(userCommunicationServiceIds))
	for userCommunicationServiceId := range userCommunicationServiceIds {
		userCommunicationServiceIdsList = append(userCommunicationServiceIdsList, userCommunicationServiceId)
	}

	// Fetch the facets from the repository
	userCommunicationServices, err := s.UserCommunicationServiceRepository.FindManyByIdsAndUserId(userCommunicationServiceIdsList, userId)
	if err != nil {
		return nil, err
	}

	if len(userCommunicationServices) != len(userCommunicationServiceIdsList) {
		return nil, fmt.Errorf("some user communication services were not found")
	}

	return userCommunicationServices, nil
}
