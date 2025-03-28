package dto

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
	"time"
)

type UpdateFacetEnrichedConfigItem struct {
	Id      uint8                     `json:"id"`
	Status  enum.FacetStatus          `json:"status"`
	Name    string                    `json:"name"`
	Value   string                    `json:"value"`
	Service enum.CommunicationService `json:"service"`
}

type UpdateFacetEnrichedConfig struct {
	Items []UpdateFacetEnrichedConfigItem `json:"items"`
}

type UpdateFacetResponse struct {
	ID            uint8                     `json:"id"`
	Color         string                    `json:"color"`
	PublicLabel   string                    `json:"publicLabel"`
	PrivateLabel  string                    `json:"privateLabel"`
	Configuration UpdateFacetEnrichedConfig `json:"configuration"`
	CreatedAt     time.Time                 `json:"createdAt"`
}

func ToUpdateFacetResponse(facet model.Facet, communicationServices []model.UserCommunicationService) UpdateFacetResponse {
	enrichedItems := make([]UpdateFacetEnrichedConfigItem, len(facet.Configuration.Items))
	for configItemIndex, configItem := range facet.Configuration.Items {
		for _, cs := range communicationServices {
			if cs.ID == configItem.Id {
				enrichedItems[configItemIndex] = UpdateFacetEnrichedConfigItem{
					Id:      configItem.Id,
					Status:  configItem.Status,
					Name:    cs.Name,
					Value:   cs.Value,
					Service: cs.Service,
				}
			}
		}
	}

	return UpdateFacetResponse{
		ID:            facet.ID,
		Color:         facet.Color,
		PublicLabel:   facet.PublicLabel,
		PrivateLabel:  facet.PrivateLabel,
		Configuration: UpdateFacetEnrichedConfig{Items: enrichedItems},
		CreatedAt:     facet.CreatedAt,
	}
}
