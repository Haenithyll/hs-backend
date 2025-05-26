package di

import (
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"
	"hs-backend/internal/service"

	"gorm.io/gorm"
)

func InitializeRequestHandler(db *gorm.DB) *handler.RequestHandler {
	requestLevelRepository := repository.NewRequestLevelRepository(db)
	requestRepository := repository.NewRequestRepository(db)
	userRepository := repository.NewUserRepository(db)

	requestService := service.NewRequestService(
		requestLevelRepository,
		requestRepository,
		userRepository,
	)

	requestHandler := handler.NewRequestHandler(requestService)

	return requestHandler
}
