package di

import (
	"hs-backend/internal/handler/prism"
	"hs-backend/internal/repository"

	"gorm.io/gorm"
)

func InitializeCreatePrismHandler(db *gorm.DB) *prism.CreatePrismHandler {
	prismRepository := repository.NewPrismRepository(db)
	facetRepository := repository.NewFacetRepository(db)
	createPrismHandler := prism.NewCreatePrismHandler(facetRepository, prismRepository)
	return createPrismHandler
}
