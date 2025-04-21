package di

import (
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"
	"hs-backend/internal/service"

	"gorm.io/gorm"
)

func InitializeRefractedFacetHandler(db *gorm.DB) *handler.RefractedFacetHandler {
	facetRepository := repository.NewFacetRepository(db)
	userCommunicationServiceRepository := repository.NewUserCommunicationServiceRepository(db)
	userRepository := repository.NewUserRepository(db)
	userPrismTrackerRepository := repository.NewUserPrismTrackerRepository(db)

	refractedFacetService := service.NewRefractedFacetService(
		facetRepository,
		userCommunicationServiceRepository,
		userRepository,
		userPrismTrackerRepository,
	)

	refractedFacetHandler := handler.NewRefractedFacetHandler(*refractedFacetService)

	return refractedFacetHandler
}
