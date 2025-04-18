package validator

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/model/enum"
	"hs-backend/internal/request"
)

func ValidateCreateUserCommunicationServiceRequest(request *request.CreateUserCommunicationServiceRequest) *domain.DomainError {
	if err := validateCommunicationServiceValue(request.Service); err != nil {
		return err
	}

	return nil
}

func ValidateUpdateUserCommunicationServiceRequest(request *request.UpdateUserCommunicationServiceRequest) *domain.DomainError {
	if request.Service == nil {
		return nil
	}

	if err := validateCommunicationServiceValue(*request.Service); err != nil {
		return err
	}

	return nil
}

func validateCommunicationServiceValue(value enum.CommunicationService) *domain.DomainError {
	if !enum.CommunicationService(value).IsValid() {
		return domain.NewDomainError(domain.ErrBadRequest, "invalid service: "+string(value))
	}

	return nil
}
