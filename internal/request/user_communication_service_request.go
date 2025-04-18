package request

import (
	"hs-backend/internal/model/enum"
)

type CreateUserCommunicationServiceRequest struct {
	Name    string                    `json:"name" binding:"required"`
	Value   string                    `json:"value" binding:"required"`
	Service enum.CommunicationService `json:"service" binding:"required"`
}

type UpdateUserCommunicationServiceRequest struct {
	UserCommunicationServiceID uint8                      `uri:"userCommunicationServiceId" json:"-" binding:"required"`
	Name                       *string                    `json:"name,omitempty"`
	Value                      *string                    `json:"value,omitempty"`
	Service                    *enum.CommunicationService `json:"service,omitempty"`
}

type DeleteUserCommunicationServiceRequest struct {
	UserCommunicationServiceID uint8 `uri:"userCommunicationServiceId" json:"-" binding:"required"`
}
