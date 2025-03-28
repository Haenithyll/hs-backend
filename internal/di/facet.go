package di

import (
	"hs-backend/internal/handler/facet"
	"hs-backend/internal/repository"

	"gorm.io/gorm"
)

func InitializeGetFacetsHandler(db *gorm.DB) *facet.GetFacetsHandler {
	facetRepository := repository.NewFacetRepository(db)
	userCommunicationServiceRepository := repository.NewUserCommunicationServiceRepository(db)
	getFacetsHandler := facet.NewGetFacetsHandler(facetRepository, userCommunicationServiceRepository)
	return getFacetsHandler
}

func InitializeCreateFacetHandler(db *gorm.DB) *facet.CreateFacetHandler {
	facetRepository := repository.NewFacetRepository(db)
	userCommunicationServiceRepository := repository.NewUserCommunicationServiceRepository(db)
	createFacetHandler := facet.NewCreateFacetHandler(facetRepository, userCommunicationServiceRepository)
	return createFacetHandler
}

func InitializeUpdateFacetHandler(db *gorm.DB) *facet.UpdateFacetHandler {
	facetRepository := repository.NewFacetRepository(db)
	userCommunicationServiceRepository := repository.NewUserCommunicationServiceRepository(db)
	updateFacetHandler := facet.NewUpdateFacetHandler(facetRepository, userCommunicationServiceRepository)
	return updateFacetHandler
}

func InitializeDeleteFacetHandler(db *gorm.DB) *facet.DeleteFacetHandler {
	facetRepository := repository.NewFacetRepository(db)
	deleteFacetHandler := facet.NewDeleteFacetHandler(facetRepository)
	return deleteFacetHandler
}
