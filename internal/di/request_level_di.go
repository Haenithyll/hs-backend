package di

import (
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"
	"hs-backend/internal/service"

	"gorm.io/gorm"
)

func InitializeRequestLevelHandler(db *gorm.DB) *handler.RequestLevelHandler {
	requestLevelRepository := repository.NewRequestLevelRepository(db)

	requestLevelService := service.NewRequestLevelService(requestLevelRepository)

	requestLevelHandler := handler.NewRequestLevelHandler(requestLevelService)

	return requestLevelHandler
}
