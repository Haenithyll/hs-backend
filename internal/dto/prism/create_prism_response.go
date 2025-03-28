package dto

import (
	"hs-backend/internal/model"
	"time"

	"github.com/google/uuid"
)

type CreatePrismFacet struct {
	ID           uint8  `json:"id"`
	Color        string `json:"color"`
	PublicLabel  string `json:"publicLabel"`
	PrivateLabel string `json:"privateLabel"`
}

type CreatePrismEnrichedUserItem struct {
	UserId uuid.UUID        `json:"userId"`
	Facet  CreatePrismFacet `json:"facet"`
}

type CreatePrismEnrichedConfig struct {
	Base  CreatePrismFacet              `json:"base"`
	Users []CreatePrismEnrichedUserItem `json:"users"`
}

type CreatePrismResponse struct {
	ID            uint8                     `json:"id"`
	Name          string                    `json:"name"`
	Configuration CreatePrismEnrichedConfig `json:"configuration"`
	CreatedAt     time.Time                 `json:"createdAt"`
}

func ToCreatePrismResponse(prism model.Prism, facets []model.Facet) CreatePrismResponse {
	getFacetById := func(id uint8, facets []model.Facet) CreatePrismFacet {
		for _, facet := range facets {
			if facet.ID == id {
				return CreatePrismFacet{
					ID:           facet.ID,
					Color:        facet.Color,
					PublicLabel:  facet.PublicLabel,
					PrivateLabel: facet.PrivateLabel,
				}
			}
		}

		return CreatePrismFacet{}
	}

	baseFacet := getFacetById(prism.Configuration.Base, facets)

	enrichedUserItems := make([]CreatePrismEnrichedUserItem, len(prism.Configuration.Users))
	for configUserItemIndex, configUserItem := range prism.Configuration.Users {
		facet := getFacetById(configUserItem.FacetId, facets)
		enrichedUserItems[configUserItemIndex] = CreatePrismEnrichedUserItem{
			UserId: configUserItem.UserId,
			Facet:  facet,
		}
	}

	return CreatePrismResponse{
		ID:            prism.ID,
		Name:          prism.Name,
		Configuration: CreatePrismEnrichedConfig{Base: baseFacet, Users: enrichedUserItems},
		CreatedAt:     prism.CreatedAt,
	}
}
