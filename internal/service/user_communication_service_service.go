package service

import (
	"errors"
	"hs-backend/internal/domain"
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
	"hs-backend/internal/repository"
	"hs-backend/internal/request"
	"hs-backend/internal/response"
	"hs-backend/internal/response/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCommunicationServiceService struct {
	FacetRepository                    repository.FacetRepository
	UserCommunicationServiceRepository repository.UserCommunicationServiceRepository
}

func NewUserCommunicationServiceService(
	facetRepository repository.FacetRepository,
	userCommunicationServiceRepository repository.UserCommunicationServiceRepository,
) *UserCommunicationServiceService {
	return &UserCommunicationServiceService{
		FacetRepository:                    facetRepository,
		UserCommunicationServiceRepository: userCommunicationServiceRepository,
	}
}

// Fetches all user communication services associated with the given user ID.
func (s *UserCommunicationServiceService) GetAllUserCommunicationServices(userId uuid.UUID) ([]response.UserCommunicationServiceResponse, *domain.DomainError) {
	userCommunicationServices, err := s.UserCommunicationServiceRepository.FindManyByUserId(userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get user communication services: "+err.Error())
	}

	return mapper.ToUserCommunicationServiceResponses(userCommunicationServices), nil
}

// Creates a new user communication service for the given user.
func (s *UserCommunicationServiceService) CreateUserCommunicationService(userId uuid.UUID, request request.CreateUserCommunicationServiceRequest) (*response.UserCommunicationServiceResponse, *domain.DomainError) {
	userCommunicationService := model.UserCommunicationService{
		UserId:  userId,
		Name:    request.Name,
		Value:   request.Value,
		Service: enum.CommunicationService(request.Service),
	}

	err := s.UserCommunicationServiceRepository.CreateOne(&userCommunicationService)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to create user communication service: "+err.Error())
	}

	response := mapper.ToUserCommunicationServiceResponse(userCommunicationService)

	return &response, nil
}

// Updates a user communication service for the given user.
func (s *UserCommunicationServiceService) UpdateUserCommunicationService(userId uuid.UUID, request request.UpdateUserCommunicationServiceRequest) (*response.UserCommunicationServiceResponse, *domain.DomainError) {
	userCommunicationService, err := s.UserCommunicationServiceRepository.FindOneByIDAndUserID(request.UserCommunicationServiceID, userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "User communication service not found")
	}

	if request.Name != nil {
		userCommunicationService.Name = *request.Name
	}
	if request.Value != nil {
		userCommunicationService.Value = *request.Value
	}
	if request.Service != nil {
		userCommunicationService.Service = *request.Service
	}

	err = s.UserCommunicationServiceRepository.UpdateOne(userCommunicationService)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to update user communication service: "+err.Error())
	}

	response := mapper.ToUserCommunicationServiceResponse(*userCommunicationService)

	return &response, nil
}

// Deletes a user communication service for the given user.
func (s *UserCommunicationServiceService) DeleteUserCommunicationService(userId uuid.UUID, userCommunicationServiceId uint8) *domain.DomainError {
	if err := s.UserCommunicationServiceRepository.DeleteOneByIDAndUserID(userCommunicationServiceId, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.NewDomainError(domain.ErrNotFound, "user communication service not found")
		}
		return domain.NewDomainError(domain.ErrInternalServerError, "failed to delete user communication service: "+err.Error())
	}

	if err := s.FacetRepository.RemoveUserCommunicationServiceFromFacets(userCommunicationServiceId); err != nil {
		return domain.NewDomainError(domain.ErrInternalServerError, "failed to apply user communication service deletion to existing facets: "+err.Error())
	}

	return nil
}
