package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRequestLevelRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	requestLevelHandler := di.InitializeRequestLevelHandler(db)

	requestLevels := rg.Group("/requests/levels")
	{
		requestLevels.GET("", requestLevelHandler.GetAll)
	}
}
