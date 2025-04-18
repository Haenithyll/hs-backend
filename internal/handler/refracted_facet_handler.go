package handler

import (
	"hs-backend/internal/service"
)

type RefractedFacetHandler struct {
	RefractedFacetService service.RefractedFacetService
}

func NewRefractedFacetHandler(refractedFacetService service.RefractedFacetService) *RefractedFacetHandler {
	return &RefractedFacetHandler{refractedFacetService}
}
