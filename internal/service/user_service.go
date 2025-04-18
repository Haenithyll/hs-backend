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

func (s *UserService) GetUserById(userId uuid.UUID) (*response.UserResponse, *domain.DomainError) {
	user, err := s.UserRepository.FindOneById(userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "user not found")
	}

	response := mapper.ToUserResponse(user)

	return response, nil
}

func (s *UserService) GetUserByEmail(email string) (*response.UserResponse, *domain.DomainError) {
	user, err := s.UserRepository.FindOneByEmail(email)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "user not found")
	}

	response := mapper.ToUserResponse(user)

	return response, nil
}
