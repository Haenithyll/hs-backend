package dto

import (
	"hs-backend/internal/model"
	"time"
)

func ToFacetResponse(f model.Facet) FacetResponse {
	return FacetResponse{
		ID:            f.ID,
		Color:         f.Color,
		PublicLabel:   f.PublicLabel,
		PrivateLabel:  f.PrivateLabel,
		Configuration: f.Configuration,
		CreatedAt:     f.CreatedAt,
	}
}

func ToFacetResponses(facets []model.Facet) []FacetResponse {
	responses := make([]FacetResponse, len(facets))
	for i, facet := range facets {
		responses[i] = ToFacetResponse(facet)
	}
	return responses
}

type FacetResponse struct {
	ID            int8              `json:"id"`
	Color         string            `json:"color"`
	PublicLabel   string            `json:"publicLabel"`
	PrivateLabel  string            `json:"privateLabel"`
	Configuration model.FacetConfig `json:"configuration"`
	CreatedAt     time.Time         `json:"createdAt"`
}
