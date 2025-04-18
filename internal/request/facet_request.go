package request

import "hs-backend/internal/model/json"

type CreateFacetRequest struct {
	Color         string           `json:"color" binding:"required"`
	PublicLabel   string           `json:"publicLabel" binding:"required"`
	PrivateLabel  string           `json:"privateLabel" binding:"required"`
	Configuration json.FacetConfig `json:"configuration" binding:"required"`
}

type UpdateFacetRequest struct {
	FacetID       uint8             `uri:"facetId" json:"-" binding:"required"`
	Color         *string           `json:"color,omitempty"`
	PublicLabel   *string           `json:"publicLabel,omitempty"`
	PrivateLabel  *string           `json:"privateLabel,omitempty"`
	Configuration *json.FacetConfig `json:"configuration,omitempty"`
}

type DeleteFacetRequest struct {
	FacetID uint8 `uri:"facetId" json:"-" binding:"required"`
}
