package di

import (
	rf "hs-backend/internal/handler/refracted_facet"
	"hs-backend/internal/repository"

	"gorm.io/gorm"
)

func InitializeGetRefractedFacetsHandler(db *gorm.DB) *rf.GetRefractedFacetsHandler {
	facetRepository := repository.NewFacetRepository(db)
	userRepository := repository.NewUserRepository(db)
	userPrismTrackerRepository := repository.NewUserPrismTrackerRepository(db)
	getRefractedFacetsHandler := rf.NewGetRefractedFacetsHandler(facetRepository, userRepository, userPrismTrackerRepository)
	return getRefractedFacetsHandler
}
