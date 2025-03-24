package dto

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/json"
	"time"
)

type GetFacetResponseItem struct {
	ID            int8             `json:"id"`
	Color         string           `json:"color"`
	PublicLabel   string           `json:"publicLabel"`
	PrivateLabel  string           `json:"privateLabel"`
	Configuration json.FacetConfig `json:"configuration"`
	CreatedAt     time.Time        `json:"createdAt"`
}

type GetFacetsResponse []GetFacetResponseItem

func ToGetFacetsResponse(facets []model.Facet) GetFacetsResponse {
	gfrItems := make([]GetFacetResponseItem, len(facets))
	for i, facet := range facets {
		gfrItems[i] = GetFacetResponseItem{
			ID:            facet.ID,
			Color:         facet.Color,
			PublicLabel:   facet.PublicLabel,
			PrivateLabel:  facet.PrivateLabel,
			Configuration: facet.Configuration,
			CreatedAt:     facet.CreatedAt,
		}
	}
	return gfrItems
}
