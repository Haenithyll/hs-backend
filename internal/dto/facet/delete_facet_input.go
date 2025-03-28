package dto

type DeleteFacetInput struct {
	FacetID uint8 `uri:"facetId" binding:"required"`
}
