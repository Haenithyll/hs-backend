package dto

import (
	"fmt"
	"hs-backend/internal/model/enum"
)

type UpdateUserCommunicationServiceInput struct {
	UserCommunicationServiceID uint8                      `uri:"userCommunicationServiceId" json:"-" binding:"required"`
	Name                       *string                    `json:"name,omitempty"`
	Value                      *string                    `json:"value,omitempty"`
	Service                    *enum.CommunicationService `json:"service,omitempty"`
}

func (u UpdateUserCommunicationServiceInput) Validate() error {
	if u.Service != nil && !enum.CommunicationService(*u.Service).IsValid() {
		return fmt.Errorf("invalid service: %s", *u.Service)
	}

	return nil
}
