package service

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/repository"
	"hs-backend/internal/response"
	"hs-backend/internal/response/mapper"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) GetAllUsers() (*response.UserResponses, *domain.DomainError) {
	users, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get all users: "+err.Error())
	}

	responses := mapper.ToUserResponses(users)

	return &responses, nil
}

func (s *UserService) GetUserById(userId uuid.UUID) (*response.UserResponse, *domain.DomainError) {
	user, err := s.UserRepository.FindOneById(userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "user not found")
	}

	response := mapper.ToUserResponse(user)

	return response, nil
}
