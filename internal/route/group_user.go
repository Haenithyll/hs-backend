package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userHandler := di.InitializeUserHandler(db)

	users := rg.Group("/users")
	{
		users.GET("/me", userHandler.GetUserMe)
		users.GET("/email/:email", userHandler.GetUserByEmail)
	}
}
