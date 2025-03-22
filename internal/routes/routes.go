package routes

import (
	"hs-backend/internal/middleware"
	"hs-backend/internal/shared"
	"hs-backend/internal/user/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	authenticated := api.Group("/")
	authenticated.Use(middleware.AuthMiddleware())

	deps := &shared.HandlerDeps{DB: db}

	users := authenticated.Group("/users")
	{
		users.GET("", handler.NewGetUserByEmailHandler(deps).Handle)
	}
}
