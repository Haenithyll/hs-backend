package di

import (
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"
	"hs-backend/internal/service"

	"gorm.io/gorm"
)

func InitializeFacetHandler(db *gorm.DB) *handler.FacetHandler {
	facetRepository := repository.NewFacetRepository(db)
	userCommunicationServiceRepository := repository.NewUserCommunicationServiceRepository(db)

	facetService := service.NewFacetService(facetRepository, userCommunicationServiceRepository)

	facetHandler := handler.NewFacetHandler(facetService)

	return facetHandler
}
