package dto

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
	"time"
)

type UpdateUserCommunicationServiceResponse struct {
	ID        uint8                     `json:"id"`
	Name      string                    `json:"name"`
	Value     string                    `json:"value"`
	Service   enum.CommunicationService `json:"service"`
	CreatedAt time.Time                 `json:"createdAt"`
}

func ToUpdateUserCommunicationServiceResponse(userCommunicationService model.UserCommunicationService) UpdateUserCommunicationServiceResponse {
	return UpdateUserCommunicationServiceResponse{
		ID:        userCommunicationService.ID,
		Name:      userCommunicationService.Name,
		Value:     userCommunicationService.Value,
		Service:   userCommunicationService.Service,
		CreatedAt: userCommunicationService.CreatedAt,
	}
}
