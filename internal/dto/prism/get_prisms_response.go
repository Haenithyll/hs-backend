package dto

import (
	"hs-backend/internal/model"
	"time"

	"github.com/google/uuid"
)

type GetPrismFacet struct {
	ID           uint8  `json:"id"`
	Color        string `json:"color"`
	PublicLabel  string `json:"publicLabel"`
	PrivateLabel string `json:"privateLabel"`
}

type GetPrismEnrichedUserItem struct {
	UserId uuid.UUID     `json:"userId"`
	Facet  GetPrismFacet `json:"facet"`
}

type GetPrismEnrichedConfig struct {
	Base  GetPrismFacet              `json:"base"`
	Users []GetPrismEnrichedUserItem `json:"users"`
}

type GetPrismResponseItem struct {
	ID            uint8                  `json:"id"`
	Name          string                 `json:"name"`
	Configuration GetPrismEnrichedConfig `json:"configuration"`
	CreatedAt     time.Time              `json:"createdAt"`
}

type GetPrismsResponse []GetPrismResponseItem

func ToGetPrismsResponse(prisms []model.Prism, facets []model.Facet) GetPrismsResponse {
	getFacetById := func(id uint8, facets []model.Facet) GetPrismFacet {
		for _, facet := range facets {
			if facet.ID == id {
				return GetPrismFacet{
					ID:           facet.ID,
					Color:        facet.Color,
					PublicLabel:  facet.PublicLabel,
					PrivateLabel: facet.PrivateLabel,
				}
			}
		}

		return GetPrismFacet{}
	}

	gprItems := make([]GetPrismResponseItem, len(prisms))

	for prismIndex, prism := range prisms {
		baseFacet := getFacetById(prism.Configuration.Base, facets)

		enrichedUserItems := make([]GetPrismEnrichedUserItem, len(prism.Configuration.Users))
		for configUserItemIndex, configUserItem := range prism.Configuration.Users {
			facet := getFacetById(configUserItem.FacetId, facets)
			enrichedUserItems[configUserItemIndex] = GetPrismEnrichedUserItem{
				UserId: configUserItem.UserId,
				Facet:  facet,
			}
		}

		gprItems[prismIndex] = GetPrismResponseItem{
			ID:            prism.ID,
			Name:          prism.Name,
			Configuration: GetPrismEnrichedConfig{Base: baseFacet, Users: enrichedUserItems},
			CreatedAt:     prism.CreatedAt,
		}
	}

	return gprItems
}
