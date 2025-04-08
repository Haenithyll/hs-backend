package types

import (
	"time"
)

type RefractedFacet struct {
	Id            uint8     `json:"id"`
	Label         string    `json:"label"`
	Color         string    `json:"color"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}
