package dto

import (
	"fmt"
	"hs-backend/internal/model/enum"
	"hs-backend/internal/model/json"
	"strings"
)

type UpdateFacetInput struct {
	FacetID       uint8             `uri:"facetId" json:"-" binding:"required"`
	Color         *string           `json:"color,omitempty"`
	PublicLabel   *string           `json:"publicLabel,omitempty"`
	PrivateLabel  *string           `json:"privateLabel,omitempty"`
	Configuration *json.FacetConfig `json:"configuration,omitempty"`
}

func (c *UpdateFacetInput) Validate() error {
	var errs []string
	for _, item := range c.Configuration.Items {
		if !enum.FacetStatus(item.Status).IsValid() {
			errs = append(errs, fmt.Sprintf("invalid status for config id %d: %s", item.Id, item.Status))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(errs, ", "))
	}
	return nil
}
