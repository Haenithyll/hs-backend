package di

import (
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"
	"hs-backend/internal/service"

	"gorm.io/gorm"
)

func InitializePrismHandler(db *gorm.DB) *handler.PrismHandler {
	prismRepository := repository.NewPrismRepository(db)
	facetRepository := repository.NewFacetRepository(db)
	userPrismTrackerRepository := repository.NewUserPrismTrackerRepository(db)
	userRepository := repository.NewUserRepository(db)

	prismService := service.NewPrismService(prismRepository, facetRepository, userPrismTrackerRepository, userRepository)

	prismHandler := handler.NewPrismHandler(prismService)

	return prismHandler
}
