package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
	"hs-backend/internal/response"
)

func ToFacetResponses(facets []model.Facet, communicationServiceMap map[uint8]model.UserCommunicationService) response.FacetResponses {
	facetResponses := make([]response.FacetResponse, len(facets))
	for i, facet := range facets {
		facetResponses[i] = ToFacetResponse(facet, communicationServiceMap)
	}
	return facetResponses
}

func ToFacetResponse(facet model.Facet, communicationServiceMap map[uint8]model.UserCommunicationService) response.FacetResponse {
	enrichedConfigItems := make([]response.FacetResponseEnrichedConfigItem, len(facet.Configuration.Items))

	for i, item := range facet.Configuration.Items {
		communicationService := newFacetResponseEnrichedConfigItemCommunicationService(communicationServiceMap[item.Id])

		enrichedConfigItems[i] = newFacetResponseEnrichedConfigItem(
			communicationService,
			item.Status,
		)
	}

	enrichedConfig := newFacetResponseEnrichedConfig(enrichedConfigItems)

	return newFacetResponse(facet, enrichedConfig)
}

func newFacetResponse(facet model.Facet, enrichedConfig response.FacetResponseEnrichedConfig) response.FacetResponse {
	return response.FacetResponse{
		ID:            facet.ID,
		Color:         facet.Color,
		PublicLabel:   facet.PublicLabel,
		PrivateLabel:  facet.PrivateLabel,
		Configuration: enrichedConfig,
		CreatedAt:     facet.CreatedAt,
	}
}

func newFacetResponseEnrichedConfig(enrichedConfigItems []response.FacetResponseEnrichedConfigItem) response.FacetResponseEnrichedConfig {
	return response.FacetResponseEnrichedConfig{Items: enrichedConfigItems}
}

func newFacetResponseEnrichedConfigItem(
	communicationService response.FacetResponseEnrichedConfigItemCommunicationService,
	status enum.FacetStatus,
) response.FacetResponseEnrichedConfigItem {
	return response.FacetResponseEnrichedConfigItem{
		CommunicationService: communicationService,
		Status:               status,
	}
}

func newFacetResponseEnrichedConfigItemCommunicationService(
	communicationService model.UserCommunicationService,
) response.FacetResponseEnrichedConfigItemCommunicationService {
	return response.FacetResponseEnrichedConfigItemCommunicationService{
		ID:      communicationService.ID,
		Name:    communicationService.Name,
		Value:   communicationService.Value,
		Service: communicationService.Service,
	}
}
