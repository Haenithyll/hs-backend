package validator

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/model/json"
	"hs-backend/internal/request"
	"hs-backend/internal/util"

	"github.com/google/uuid"
)

func ValidateCreatePrismRequest(r *request.CreatePrismRequest) *domain.DomainError {
	if err := validateConfigUserIdsAreUnique(r.Configuration); err != nil {
		return err
	}

	return nil
}

func ValidateUpdatePrismRequest(r *request.UpdatePrismRequest) *domain.DomainError {
	if r.Configuration == nil {
		return nil
	}

	if err := validateConfigUserIdsAreUnique(*r.Configuration); err != nil {
		return err
	}

	return nil
}

func validateConfigUserIdsAreUnique(configuration json.PrismConfig) *domain.DomainError {
	userIds := make([]uuid.UUID, len(configuration.Users))

	for index, user := range configuration.Users {
		userIds[index] = user.UserId
	}

	if !util.AreUUIDsUnique(userIds) {
		return domain.NewDomainError(domain.ErrBadRequest, "user ids must be unique")
	}

	return nil
}
