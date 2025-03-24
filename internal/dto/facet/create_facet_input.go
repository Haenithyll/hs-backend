package dto

import (
	"fmt"
	"hs-backend/internal/model/enum"
	"hs-backend/internal/model/json"
	"strings"
)

type CreateFacetInput struct {
	Color         string           `json:"color" binding:"required"`
	PublicLabel   string           `json:"publicLabel" binding:"required"`
	PrivateLabel  string           `json:"privateLabel" binding:"required"`
	Configuration json.FacetConfig `json:"configuration" binding:"required"`
}

func (c *CreateFacetInput) Validate() error {
	var errs []string
	for _, item := range c.Configuration.Items {
		if !enum.FacetStatus(item.Status).IsValid() {
			errs = append(errs, fmt.Sprintf("invalid status for config id %d: %s", item.Id, item.Status))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation errors:\n%s", strings.Join(errs, "\n"))
	}
	return nil
}
