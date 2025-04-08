package dto

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
	"time"
)

type GetFacetEnrichedConfigItem struct {
	Id      uint8                     `json:"id"`
	Status  enum.FacetStatus          `json:"status"`
	Name    string                    `json:"name"`
	Value   string                    `json:"value"`
	Service enum.CommunicationService `json:"service"`
}

type GetFacetEnrichedConfig struct {
	Items []GetFacetEnrichedConfigItem `json:"items"`
}

type GetFacetResponseItem struct {
	ID            uint8                  `json:"id"`
	Color         string                 `json:"color"`
	PublicLabel   string                 `json:"publicLabel"`
	PrivateLabel  string                 `json:"privateLabel"`
	Configuration GetFacetEnrichedConfig `json:"configuration"`
	CreatedAt     time.Time              `json:"createdAt"`
}

type GetFacetsResponse []GetFacetResponseItem

func ToGetFacetsResponse(facets []model.Facet, communicationServices []model.UserCommunicationService) GetFacetsResponse {
	gfrItems := make([]GetFacetResponseItem, len(facets))

	for facetIndex, facet := range facets {
		enrichedItems := make([]GetFacetEnrichedConfigItem, len(facet.Configuration.Items))

		for configItemIndex, configItem := range facet.Configuration.Items {
			for _, cs := range communicationServices {
				if cs.ID == configItem.Id {
					enrichedItems[configItemIndex] = GetFacetEnrichedConfigItem{
						Id:      configItem.Id,
						Status:  configItem.Status,
						Name:    cs.Name,
						Value:   cs.Value,
						Service: cs.Service,
					}
				}
			}
		}
		gfrItems[facetIndex] = GetFacetResponseItem{
			ID:            facet.ID,
			Color:         facet.Color,
			PublicLabel:   facet.PublicLabel,
			PrivateLabel:  facet.PrivateLabel,
			Configuration: GetFacetEnrichedConfig{Items: enrichedItems},
			CreatedAt:     facet.CreatedAt,
		}
	}
	return gfrItems
}
