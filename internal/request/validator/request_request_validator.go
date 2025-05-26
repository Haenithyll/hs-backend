package validator

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/request"

	"github.com/google/uuid"
)

func ValidateCreateRequestRequest(r *request.CreateRequestRequest) *domain.DomainError {
	_, err := uuid.Parse(r.ReceiverID)
	if err != nil {
		return domain.NewDomainError(domain.ErrBadRequest, "receiver id is invalid")
	}

	if len(r.Topic) == 0 {
		return domain.NewDomainError(domain.ErrBadRequest, "topic is required")
	}

	if len(r.Topic) > 140 {
		return domain.NewDomainError(domain.ErrBadRequest, "topic must be less than 140 characters")
	}

	return nil
}
