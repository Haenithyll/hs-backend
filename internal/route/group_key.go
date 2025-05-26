package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterKeyRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	keyHandler := di.InitializeKeyHandler(db)

	keys := rg.Group("/key")
	{
		keys.POST("/generate", keyHandler.Generate)
	}
}
