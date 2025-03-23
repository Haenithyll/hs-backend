package dto

import (
	"fmt"
	"hs-backend/internal/model"
)

type CreateFacetInput struct {
	Color         string            `json:"color" binding:"required"`
	PublicLabel   string            `json:"publicLabel" binding:"required"`
	PrivateLabel  string            `json:"privateLabel" binding:"required"`
	Configuration model.FacetConfig `json:"configuration" binding:"required"`
}

var allowedStatuses = map[model.FacetStatus]bool{
	model.Available:     true,
	model.EmergencyOnly: true,
}

func (c *CreateFacetInput) Validate() error {
	for _, item := range c.Configuration.Items {
		if !allowedStatuses[item.Status] {
			return fmt.Errorf("invalid status for config key %d: %s", item.Id, item.Status)
		}
	}
	return nil
}
