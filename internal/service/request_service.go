package service

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/model"
	"hs-backend/internal/repository"
	"hs-backend/internal/request"
	"hs-backend/internal/response"
	"hs-backend/internal/response/mapper"

	"github.com/google/uuid"
)

type RequestService struct {
	RequestLevelRepository repository.RequestLevelRepository
	RequestRepository      repository.RequestRepository
	UserRepository         repository.UserRepository
}

func NewRequestService(
	requestLevelRepository repository.RequestLevelRepository,
	requestRepository repository.RequestRepository,
	userRepository repository.UserRepository,
) *RequestService {
	return &RequestService{
		RequestLevelRepository: requestLevelRepository,
		RequestRepository:      requestRepository,
		UserRepository:         userRepository,
	}
}

// Gets all received requests for the user.
func (s *RequestService) GetReceivedRequests(userId uuid.UUID) (*response.RequestResponses, *domain.DomainError) {
	enrichedRequests, err := s.RequestRepository.FindManyEnrichedByReceiverId(userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "failed to get received requests: "+err.Error())
	}

	responses := mapper.ToRequestResponsesFromEnrichedRequests(enrichedRequests)

	return &responses, nil
}

// Gets all issued requests for the user.
func (s *RequestService) GetIssuedRequests(userId uuid.UUID) (*response.RequestResponses, *domain.DomainError) {
	enrichedRequests, err := s.RequestRepository.FindManyEnrichedByIssuerId(userId)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "failed to get issued requests: "+err.Error())
	}

	responses := mapper.ToRequestResponsesFromEnrichedRequests(enrichedRequests)

	return &responses, nil
}

// Creates a new request.
// It ensures that the receiver exists and the level exists.
// It also ensures that the user is not sending a request to themselves.
func (s *RequestService) CreateRequest(userId uuid.UUID, request request.CreateRequestRequest) (*response.RequestResponse, *domain.DomainError) {
	if request.ReceiverID == userId.String() {
		return nil, domain.NewDomainError(domain.ErrBadRequest, "cannot send request to yourself")
	}

	receiverExistsById, err := s.UserRepository.ExistsById(uuid.MustParse(request.ReceiverID))
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to check if receiver exists: "+err.Error())
	}
	if !receiverExistsById {
		return nil, domain.NewDomainError(domain.ErrNotFound, "receiver not found")
	}

	levelExistsById, err := s.RequestLevelRepository.ExistsById(request.Level)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to check if level exists: "+err.Error())
	}
	if !levelExistsById {
		return nil, domain.NewDomainError(domain.ErrNotFound, "level not found")
	}

	newRequest := model.Request{
		IssuerId:   userId,
		ReceiverId: uuid.MustParse(request.ReceiverID),
		Topic:      request.Topic,
		LevelId:    request.Level,
	}

	err = s.RequestRepository.CreateOne(&newRequest)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "failed to create request: "+err.Error())
	}

	enrichedRequest, err := s.RequestRepository.FindOneEnrichedById(newRequest.ID)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrNotFound, "request not found: "+err.Error())
	}

	return mapper.ToRequestResponseFromEnrichedRequest(enrichedRequest), nil
}

// Marks a request as read if it exists and the user is the receiver.
// Ignores the request if it is already read.
func (s *RequestService) MarkRequestAsRead(userId uuid.UUID, request request.MarkRequestAsReadRequest) *domain.DomainError {
	requestModel, err := s.RequestRepository.FindOneByIdAndReceiverId(request.RequestID, userId)
	if err != nil {
		return domain.NewDomainError(domain.ErrNotFound, "request not found")
	}

	if requestModel.IsRead {
		return nil
	}

	err = s.RequestRepository.MarkOneAsRead(requestModel.ID)
	if err != nil {
		return domain.NewDomainError(domain.ErrInternalServerError, "failed to mark request as read: "+err.Error())
	}

	return nil
}
