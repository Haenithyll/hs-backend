package dto

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
	"time"
)

type CreateFacetEnrichedConfigItem struct {
	Id      uint8                     `json:"id"`
	Status  enum.FacetStatus          `json:"status"`
	Name    string                    `json:"name"`
	Value   string                    `json:"value"`
	Service enum.CommunicationService `json:"service"`
}

type CreateFacetEnrichedConfig struct {
	Items []CreateFacetEnrichedConfigItem `json:"items"`
}

type CreateFacetResponse struct {
	ID            uint8                     `json:"id"`
	Color         string                    `json:"color"`
	PublicLabel   string                    `json:"publicLabel"`
	PrivateLabel  string                    `json:"privateLabel"`
	Configuration CreateFacetEnrichedConfig `json:"configuration"`
	CreatedAt     time.Time                 `json:"createdAt"`
}

func ToCreateFacetResponse(facet model.Facet, communicationServices []model.UserCommunicationService) CreateFacetResponse {
	enrichedItems := make([]CreateFacetEnrichedConfigItem, len(facet.Configuration.Items))
	for configItemIndex, configItem := range facet.Configuration.Items {
		for _, cs := range communicationServices {
			if cs.ID == configItem.Id {
				enrichedItems[configItemIndex] = CreateFacetEnrichedConfigItem{
					Id:      configItem.Id,
					Status:  configItem.Status,
					Name:    cs.Name,
					Value:   cs.Value,
					Service: cs.Service,
				}
			}
		}
	}

	return CreateFacetResponse{
		ID:            facet.ID,
		Color:         facet.Color,
		PublicLabel:   facet.PublicLabel,
		PrivateLabel:  facet.PrivateLabel,
		Configuration: CreateFacetEnrichedConfig{Items: enrichedItems},
		CreatedAt:     facet.CreatedAt,
	}
}
