package dto

import (
	"hs-backend/internal/model"
	"time"

	"github.com/google/uuid"
)

type UpdatePrismFacet struct {
	ID           uint8  `json:"id"`
	Color        string `json:"color"`
	PublicLabel  string `json:"publicLabel"`
	PrivateLabel string `json:"privateLabel"`
}

type UpdatePrismEnrichedUserItem struct {
	UserId uuid.UUID        `json:"userId"`
	Facet  UpdatePrismFacet `json:"facet"`
}

type UpdatePrismEnrichedConfig struct {
	Base  UpdatePrismFacet              `json:"base"`
	Users []UpdatePrismEnrichedUserItem `json:"users"`
}

type UpdatePrismResponse struct {
	ID            uint8                     `json:"id"`
	Name          string                    `json:"name"`
	Configuration UpdatePrismEnrichedConfig `json:"configuration"`
	CreatedAt     time.Time                 `json:"createdAt"`
}

func ToUpdatePrismResponse(prism model.Prism, facets []model.Facet) UpdatePrismResponse {
	getFacetById := func(id uint8, facets []model.Facet) UpdatePrismFacet {
		for _, facet := range facets {
			if facet.ID == id {
				return UpdatePrismFacet{
					ID:           facet.ID,
					Color:        facet.Color,
					PublicLabel:  facet.PublicLabel,
					PrivateLabel: facet.PrivateLabel,
				}
			}
		}

		return UpdatePrismFacet{}
	}

	baseFacet := getFacetById(prism.Configuration.Base, facets)

	enrichedUserItems := make([]UpdatePrismEnrichedUserItem, len(prism.Configuration.Users))
	for configUserItemIndex, configUserItem := range prism.Configuration.Users {
		facet := getFacetById(configUserItem.FacetId, facets)
		enrichedUserItems[configUserItemIndex] = UpdatePrismEnrichedUserItem{
			UserId: configUserItem.UserId,
			Facet:  facet,
		}
	}

	return UpdatePrismResponse{
		ID:            prism.ID,
		Name:          prism.Name,
		Configuration: UpdatePrismEnrichedConfig{Base: baseFacet, Users: enrichedUserItems},
		CreatedAt:     prism.CreatedAt,
	}
}
