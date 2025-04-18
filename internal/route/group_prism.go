package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPrismRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	prismHandler := di.InitializePrismHandler(db)

	prisms := rg.Group("/prisms")
	{
		prisms.GET("", prismHandler.GetAll)
		prisms.POST("", prismHandler.Create)
		prisms.PATCH("/:prismId", prismHandler.Update)
		prisms.DELETE("/:prismId", prismHandler.Delete)
	}
}
