package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/response"

	"github.com/google/uuid"
)

func ToPrismResponses(
	prisms []model.Prism,
	facetsMap map[uint8]model.Facet,
	usersMap map[uuid.UUID]model.User,
	activePrismId *uint8,
) response.PrismResponses {
	prismResponses := make([]response.PrismResponse, len(prisms))

	for i, prism := range prisms {
		isActive := activePrismId != nil && *activePrismId == prism.ID
		prismResponses[i] = ToPrismResponse(prism, facetsMap, usersMap, isActive)
	}

	return prismResponses
}

func ToPrismResponse(
	prism model.Prism,
	facetsMap map[uint8]model.Facet,
	usersMap map[uuid.UUID]model.User,
	isActive bool,
) response.PrismResponse {
	baseFacet := newPrismResponseFacet(facetsMap[prism.Configuration.Base])

	enrichedUserItems := make([]response.PrismResponseEnrichedUserItem, len(prism.Configuration.Users))
	for i, configUserItem := range prism.Configuration.Users {

		user := newPrismResponseUser(usersMap[configUserItem.UserId])
		facet := newPrismResponseFacet(facetsMap[configUserItem.FacetId])

		enrichedUserItems[i] = newPrismResponseEnrichedUserItem(user, facet)
	}

	return newPrismResponse(prism, baseFacet, enrichedUserItems, isActive)
}

func newPrismResponse(
	prism model.Prism,
	baseFacet response.PrismResponseFacet,
	enrichedUserItems []response.PrismResponseEnrichedUserItem,
	isActive bool,
) response.PrismResponse {
	return response.PrismResponse{
		ID:            prism.ID,
		Name:          prism.Name,
		Configuration: response.PrismResponseEnrichedConfig{Base: baseFacet, Users: enrichedUserItems},
		CreatedAt:     prism.CreatedAt,
		IsActive:      isActive,
	}
}

func newPrismResponseFacet(facet model.Facet) response.PrismResponseFacet {
	return response.PrismResponseFacet{
		ID:           facet.ID,
		Color:        facet.Color,
		PublicLabel:  facet.PublicLabel,
		PrivateLabel: facet.PrivateLabel,
	}
}

func newPrismResponseUser(user model.User) response.PrismResponseUser {
	return response.PrismResponseUser{
		ID:        user.ID,
		AvatarUrl: user.AvatarURL,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func newPrismResponseEnrichedUserItem(user response.PrismResponseUser, facet response.PrismResponseFacet) response.PrismResponseEnrichedUserItem {
	return response.PrismResponseEnrichedUserItem{
		User:  user,
		Facet: facet,
	}
}
