package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/response"

	"github.com/google/uuid"
)

func ToPrismResponses(prisms []model.Prism, facetsMap map[uint8]model.Facet, usersMap map[uuid.UUID]model.User) response.PrismResponses {
	prismResponses := make([]response.PrismResponse, len(prisms))

	for i, prism := range prisms {
		prismResponses[i] = ToPrismResponse(prism, facetsMap, usersMap)
	}

	return prismResponses
}

func ToPrismResponse(prism model.Prism, facetsMap map[uint8]model.Facet, usersMap map[uuid.UUID]model.User) response.PrismResponse {
	baseFacet := newPrismResponseFacet(facetsMap[prism.Configuration.Base])

	enrichedUserItems := make([]response.PrismResponseEnrichedUserItem, len(prism.Configuration.Users))
	for i, configUserItem := range prism.Configuration.Users {

		user := newPrismResponseUser(usersMap[configUserItem.UserId])
		facet := newPrismResponseFacet(facetsMap[configUserItem.FacetId])

		enrichedUserItems[i] = newPrismResponseEnrichedUserItem(user, facet)
	}

	return newPrismResponse(prism, baseFacet, enrichedUserItems)
}

func newPrismResponse(prism model.Prism, baseFacet response.PrismResponseFacet, enrichedUserItems []response.PrismResponseEnrichedUserItem) response.PrismResponse {
	return response.PrismResponse{
		ID:            prism.ID,
		Name:          prism.Name,
		Configuration: response.PrismResponseEnrichedConfig{Base: baseFacet, Users: enrichedUserItems},
		CreatedAt:     prism.CreatedAt,
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
