package service

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/repository"
	"hs-backend/internal/response"
	"hs-backend/internal/response/mapper"
)

type RequestLevelService struct {
	RequestLevelRepository repository.RequestLevelRepository
}

func NewRequestLevelService(
	requestLevelRepository repository.RequestLevelRepository,
) *RequestLevelService {
	return &RequestLevelService{
		RequestLevelRepository: requestLevelRepository,
	}
}

// Gets all request levels.
func (s *RequestLevelService) GetAllRequestLevels() (*response.RequestLevelResponses, *domain.DomainError) {
	requestLevels, err := s.RequestLevelRepository.FindAll()
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get request levels: "+err.Error())
	}

	if len(requestLevels) == 0 {
		return &response.RequestLevelResponses{}, nil
	}

	responses := mapper.ToRequestLevelResponses(requestLevels)

	return &responses, nil
}
