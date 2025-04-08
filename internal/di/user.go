package di

import (
	"hs-backend/internal/handler/user"
	"hs-backend/internal/repository"

	"gorm.io/gorm"
)

func InitializeGetUserMeHandler(db *gorm.DB) *user.GetUserMeHandler {
	userRepository := repository.NewUserRepository(db)
	getUserMeHandler := user.NewGetUserMeHandler(userRepository)
	return getUserMeHandler
}

func InitializeGetUserByEmailHandler(db *gorm.DB) *user.GetUserByEmailHandler {
	userRepository := repository.NewUserRepository(db)
	getUserByEmailHandler := user.NewGetUserByEmailHandler(userRepository)
	return getUserByEmailHandler
}
