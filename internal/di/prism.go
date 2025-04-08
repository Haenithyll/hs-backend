package di

import (
	"hs-backend/internal/handler/prism"
	"hs-backend/internal/repository"

	"gorm.io/gorm"
)

func InitializeGetPrismsHandler(db *gorm.DB) *prism.GetPrismsHandler {
	prismRepository := repository.NewPrismRepository(db)
	facetRepository := repository.NewFacetRepository(db)
	getPrismsHandler := prism.NewGetPrismsHandler(facetRepository, prismRepository)
	return getPrismsHandler
}

func InitializeCreatePrismHandler(db *gorm.DB) *prism.CreatePrismHandler {
	prismRepository := repository.NewPrismRepository(db)
	facetRepository := repository.NewFacetRepository(db)
	createPrismHandler := prism.NewCreatePrismHandler(facetRepository, prismRepository)
	return createPrismHandler
}

func InitializeActivatePrismHandler(db *gorm.DB) *prism.ActivatePrismHandler {
	prismRepository := repository.NewPrismRepository(db)
	userPrismTrackerRepository := repository.NewUserPrismTrackerRepository(db)
	activatePrismHandler := prism.NewActivatePrismHandler(prismRepository, userPrismTrackerRepository)
	return activatePrismHandler
}

func InitializeUpdatePrismHandler(db *gorm.DB) *prism.UpdatePrismHandler {
	prismRepository := repository.NewPrismRepository(db)
	facetRepository := repository.NewFacetRepository(db)
	updatePrismHandler := prism.NewUpdatePrismHandler(facetRepository, prismRepository)
	return updatePrismHandler
}

func InitializeDeletePrismHandler(db *gorm.DB) *prism.DeletePrismHandler {
	prismRepository := repository.NewPrismRepository(db)
	deletePrismHandler := prism.NewDeletePrismHandler(prismRepository)
	return deletePrismHandler
}
