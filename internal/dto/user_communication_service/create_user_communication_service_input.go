package dto

import (
	"fmt"
	"hs-backend/internal/model/enum"
)

type CreateUserCommunicationServiceInput struct {
	Name    string                    `json:"name" binding:"required"`
	Value   string                    `json:"value" binding:"required"`
	Service enum.CommunicationService `json:"service" binding:"required"`
}

func (c *CreateUserCommunicationServiceInput) Validate() error {
	if !enum.CommunicationService(c.Service).IsValid() {
		return fmt.Errorf("invalid service: %s", c.Service)
	}

	return nil
}
