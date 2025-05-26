package validator

import (
	"fmt"
	"hs-backend/internal/domain"
	"hs-backend/internal/model/enum"
	"hs-backend/internal/model/json"
	"hs-backend/internal/request"
	"hs-backend/internal/util"
	"strings"
)

func ValidateCreateFacetRequest(request request.CreateFacetRequest) *domain.DomainError {
	err := validateFacetConfig(request.Configuration)
	if err != nil {
		return err
	}

	return nil
}

func ValidateUpdateFacetRequest(request request.UpdateFacetRequest) *domain.DomainError {
	if request.Configuration == nil {
		return nil
	}

	err := validateFacetConfig(*request.Configuration)
	if err != nil {
		return err
	}

	return nil
}

func validateFacetConfig(config json.FacetConfig) *domain.DomainError {
	errs := validateFacetStatusValues(config.Items)
	errs = append(errs, validateFacetUserCommunicationServiceIdsAreUnique(config.Items)...)

	if len(errs) > 0 {
		return domain.NewDomainError(domain.ErrBadRequest, strings.Join(errs, ", "))
	}

	return nil
}

func validateFacetStatusValues(facetConfigItems []json.FacetConfigItem) []string {
	errs := make([]string, 0)
	for _, facetConfigItem := range facetConfigItems {
		if err := validateFacetStatusValue(facetConfigItem.Status); err != nil {
			errs = append(errs, err.Error())
		}
	}
	return errs
}

func validateFacetStatusValue(status enum.FacetStatus) error {
	if !enum.FacetStatus(status).IsValid() {
		return fmt.Errorf("invalid status: %s", status)
	}
	return nil
}

func validateFacetUserCommunicationServiceIdsAreUnique(facetConfigItems []json.FacetConfigItem) []string {
	errs := make([]string, 0)
	userCommunicationServiceIds := make([]uint8, len(facetConfigItems))

	for index, item := range facetConfigItems {
		userCommunicationServiceIds[index] = item.Id
	}

	if !util.AreUintsUnique(userCommunicationServiceIds) {
		errs = append(errs, "user communication service ids must be unique")
	}

	return errs
}
