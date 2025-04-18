package validation

import (
	"fmt"
	"hs-backend/internal/domain"
	"hs-backend/internal/model"
	"hs-backend/internal/model/json"
)

// Checks if the user communication service IDs in the facet configuration are valid.
func ValidateUserCommunicationServiceIds(configuration json.FacetConfig, userCommunicationServices []model.UserCommunicationService) *domain.DomainError {
	validUserCommunicationServiceIDs := make(map[uint8]bool)

	for _, userCommunicationService := range userCommunicationServices {
		validUserCommunicationServiceIDs[userCommunicationService.ID] = true
	}

	// Anonymous function to check if a user communication service ID is present in the user communication services list.
	isUserCommunicationServiceValid := func(userCommunicationServiceId uint8) error {
		if !validUserCommunicationServiceIDs[userCommunicationServiceId] {
			return fmt.Errorf("user communication service with ID %d not found", userCommunicationServiceId)
		}

		return nil
	}

	for _, item := range configuration.Items {
		if err := isUserCommunicationServiceValid(item.Id); err != nil {
			return domain.NewDomainError(domain.ErrBadRequest, err.Error())
		}
	}

	return nil
}
