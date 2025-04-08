package dto

type ActivatePrismInput struct {
	PrismID uint8 `uri:"prismId" binding:"required"`
}
