package dto

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
)

type GetUserCommunicationServicesResponseItem struct {
	ID      uint8                     `json:"id"`
	Name    string                    `json:"name"`
	Value   string                    `json:"value"`
	Service enum.CommunicationService `json:"service"`
}

type GetUserCommunicationServicesResponse []GetUserCommunicationServicesResponseItem

func ToGetUserCommunicationServicesResponse(ucsList []model.UserCommunicationService) GetUserCommunicationServicesResponse {
	responseItems := make([]GetUserCommunicationServicesResponseItem, len(ucsList))
	for i, ucs := range ucsList {
		responseItems[i] = GetUserCommunicationServicesResponseItem{
			ID:      ucs.ID,
			Name:    ucs.Name,
			Value:   ucs.Value,
			Service: ucs.Service,
		}
	}
	return responseItems
}
