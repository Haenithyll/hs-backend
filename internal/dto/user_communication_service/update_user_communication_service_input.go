package dto

import (
	"fmt"
	"hs-backend/internal/model/enum"
)

type UpdateUserCommunicationServiceInput struct {
	UserCommunicationServiceID uint8                      `uri:"userCommunicationServiceId" json:"-" binding:"required"`
	Name                       *string                    `json:"name,omitempty"`
	Value                      *string                    `json:"value,omitempty"`
	ServiceType                *enum.CommunicationService `json:"serviceType,omitempty"`
}

func (u UpdateUserCommunicationServiceInput) Validate() error {
	if u.ServiceType != nil && !enum.CommunicationService(*u.ServiceType).IsValid() {
		return fmt.Errorf("invalid service: %s", *u.ServiceType)
	}

	return nil
}
