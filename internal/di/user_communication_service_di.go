package di

import (
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"
	"hs-backend/internal/service"

	"gorm.io/gorm"
)

func InitializeUserCommunicationServiceHandler(db *gorm.DB) *handler.UserCommunicationServiceHandler {
	facetRepository := repository.NewFacetRepository(db)
	userCommunicationServiceRepository := repository.NewUserCommunicationServiceRepository(db)

	userCommunicationServiceService := service.NewUserCommunicationServiceService(facetRepository, userCommunicationServiceRepository)

	userCommunicationServiceHandler := handler.NewUserCommunicationServiceHandler(userCommunicationServiceService)

	return userCommunicationServiceHandler
}
