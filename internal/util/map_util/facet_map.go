package map_util

import (
	"hs-backend/internal/model"
)

func BuildFacetMapById(facets []model.Facet) map[uint8]model.Facet {
	m := make(map[uint8]model.Facet, len(facets))
	for _, facet := range facets {
		m[facet.ID] = facet
	}
	return m
}
