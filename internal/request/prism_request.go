package request

import (
	"hs-backend/internal/model/json"
)

type CreatePrismRequest struct {
	Name          string           `json:"name" binding:"required"`
	Configuration json.PrismConfig `json:"configuration" binding:"required"`
}

type UpdatePrismRequest struct {
	PrismID       uint8             `uri:"prismId" json:"-" binding:"required"`
	Name          *string           `json:"name,omitempty"`
	Configuration *json.PrismConfig `json:"configuration,omitempty"`
}

type DeletePrismRequest struct {
	PrismID uint8 `uri:"prismId" json:"-" binding:"required"`
}

type ActivatePrismRequest struct {
	PrismID uint8 `uri:"prismId" json:"-" binding:"required"`
}
