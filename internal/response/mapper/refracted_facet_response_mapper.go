package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/enum"
	"hs-backend/internal/response"
	"time"

	"github.com/google/uuid"
)

func ToRefractedFacetResponses(
	users []model.User,
	facetByUserIdMap map[uuid.UUID]model.Facet,
	userCommunicationServicesByUserIdMap map[uuid.UUID][]model.UserCommunicationService,
	lastUpdatedAtByUserIdMap map[uuid.UUID]time.Time,
) []response.RefractedFacetResponse {
	refractedFacetResponses := make([]response.RefractedFacetResponse, 0, len(users))
	for _, user := range users {
		facet, found := facetByUserIdMap[user.ID]

		if !found {
			refractedFacetResponses = append(refractedFacetResponses, ToEmptyRefractedFacetResponse(user))
			continue
		}

		userCommunicationServices := userCommunicationServicesByUserIdMap[user.ID]
		lastUpdatedAt := lastUpdatedAtByUserIdMap[user.ID]

		refractedFacetResponses = append(refractedFacetResponses, ToRefractedFacetResponse(
			&user,
			&facet,
			userCommunicationServices,
			lastUpdatedAt,
		))
	}
	return refractedFacetResponses
}

func ToRefractedFacetResponse(
	user *model.User,
	facet *model.Facet,
	userCommunicationServices []model.UserCommunicationService,
	lastUpdatedAt time.Time,
) response.RefractedFacetResponse {
	enrichedCommunicationServiceByIdMap := make(map[uint8]response.RefractedFacetResponseEnrichedConfigItemCommunicationService)
	for _, userCommunicationService := range userCommunicationServices {
		enrichedCommunicationServiceByIdMap[userCommunicationService.ID] =
			newRefractedFacetResponseEnrichedConfigCommunicationService(userCommunicationService)
	}

	enrichedConfigItems := make([]response.RefractedFacetResponseEnrichedConfigItem, 0, len(facet.Configuration.Items))
	for _, item := range facet.Configuration.Items {
		enrichedConfigItem := newRefractedFacetResponseEnrichedConfigItem(
			enrichedCommunicationServiceByIdMap[item.Id],
			item.Status,
		)
		enrichedConfigItems = append(enrichedConfigItems, enrichedConfigItem)
	}

	enrichedConfig := newRefractedFacetResponseEnrichedConfig(enrichedConfigItems)

	refractedUser := newRefractedFacetResponseUser(*user)
	refractedFacet := newRefractedFacetResponseRefractedFacet(facet, enrichedConfig, lastUpdatedAt)

	return newRefractedFacetResponse(refractedUser, refractedFacet)
}

func ToEmptyRefractedFacetResponse(user model.User) response.RefractedFacetResponse {
	refractedResponseUser := newRefractedFacetResponseUser(user)

	return newRefractedFacetResponse(refractedResponseUser, nil)
}

func newRefractedFacetResponse(refractedUser response.RefractedFacetResponseUser, refractedFacet *response.RefractedFacet) response.RefractedFacetResponse {
	return response.RefractedFacetResponse{
		User:           refractedUser,
		RefractedFacet: refractedFacet,
	}
}

func newRefractedFacetResponseUser(user model.User) response.RefractedFacetResponseUser {
	return response.RefractedFacetResponseUser{
		ID:        user.ID,
		AvatarUrl: user.AvatarURL,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func newRefractedFacetResponseRefractedFacet(
	facet *model.Facet,
	configuration *response.RefractedFacetResponseEnrichedConfig,
	lastUpdatedAt time.Time,
) *response.RefractedFacet {
	return &response.RefractedFacet{
		Id:            facet.ID,
		Label:         facet.PublicLabel,
		Color:         facet.Color,
		Configuration: configuration,
		LastUpdatedAt: lastUpdatedAt,
	}
}

func newRefractedFacetResponseEnrichedConfig(items []response.RefractedFacetResponseEnrichedConfigItem) *response.RefractedFacetResponseEnrichedConfig {
	return &response.RefractedFacetResponseEnrichedConfig{
		Items: items,
	}
}

func newRefractedFacetResponseEnrichedConfigItem(
	communicationService response.RefractedFacetResponseEnrichedConfigItemCommunicationService,
	status enum.FacetStatus,
) response.RefractedFacetResponseEnrichedConfigItem {
	return response.RefractedFacetResponseEnrichedConfigItem{
		CommunicationService: communicationService,
		Status:               status,
	}
}

func newRefractedFacetResponseEnrichedConfigCommunicationService(
	communicationService model.UserCommunicationService,
) response.RefractedFacetResponseEnrichedConfigItemCommunicationService {
	return response.RefractedFacetResponseEnrichedConfigItemCommunicationService{
		ID:      communicationService.ID,
		Name:    communicationService.Name,
		Value:   communicationService.Value,
		Service: communicationService.Service,
	}
}
