package dto

type DeletePrismInput struct {
	PrismID uint8 `uri:"prismId" binding:"required"`
}
