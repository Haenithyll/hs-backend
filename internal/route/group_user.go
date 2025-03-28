package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	users := rg.Group("/users")
	{
		users.GET("", di.InitializeGetUserMeHandler(db).Handle)
		users.GET("/:userId", di.InitializeGetUserByEmailHandler(db).Handle)
	}
}
