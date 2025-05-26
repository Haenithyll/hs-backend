package di

import (
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"
	"hs-backend/internal/service"

	"gorm.io/gorm"
)

func InitializeKeyHandler(db *gorm.DB) *handler.KeyHandler {
	userRepository := repository.NewUserRepository(db)

	keyService := service.NewKeyService(userRepository)

	keyHandler := handler.NewKeyHandler(keyService)

	return keyHandler
}
