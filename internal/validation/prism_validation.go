package validation

import (
	"fmt"
	"hs-backend/internal/domain"
	"hs-backend/internal/model"
	"hs-backend/internal/model/json"
)

// Checks if the facet IDs in the prism configuration are valid.
func ValidateConfigurationFacetIds(configuration json.PrismConfig, facets []model.Facet) *domain.DomainError {
	validFacetIDs := make(map[uint8]bool)

	for _, facet := range facets {
		validFacetIDs[facet.ID] = true
	}

	// Anonymous function to check if a facet ID is present in the facets list.
	isFacetValid := func(facetId uint8) error {
		if !validFacetIDs[facetId] {
			return fmt.Errorf("facet with ID %d not found", facetId)
		}

		return nil
	}

	if err := isFacetValid(configuration.Base); err != nil {
		return domain.NewDomainError(domain.ErrBadRequest, err.Error())
	}

	for _, item := range configuration.Users {
		if err := isFacetValid(item.FacetId); err != nil {
			return domain.NewDomainError(domain.ErrBadRequest, err.Error())
		}
	}

	return nil
}
