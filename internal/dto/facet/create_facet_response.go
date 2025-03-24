package dto

import (
	"hs-backend/internal/model"
	"hs-backend/internal/model/json"
	"time"
)

type CreateFacetResponse struct {
	ID            int8             `json:"id"`
	Color         string           `json:"color"`
	PublicLabel   string           `json:"publicLabel"`
	PrivateLabel  string           `json:"privateLabel"`
	Configuration json.FacetConfig `json:"configuration"`
	CreatedAt     time.Time        `json:"createdAt"`
}

func ToCreateFacetResponse(facet model.Facet) CreateFacetResponse {
	return CreateFacetResponse{
		ID:            facet.ID,
		Color:         facet.Color,
		PublicLabel:   facet.PublicLabel,
		PrivateLabel:  facet.PrivateLabel,
		Configuration: facet.Configuration,
		CreatedAt:     facet.CreatedAt,
	}
}
