package di

import (
	"hs-backend/internal/handler"
	"hs-backend/internal/repository"
	"hs-backend/internal/service"

	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB) *handler.UserHandler {
	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	return userHandler
}
