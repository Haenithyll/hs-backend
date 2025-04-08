package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPrismRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	prisms := rg.Group("/prisms")
	{
		prisms.GET("", di.InitializeGetPrismsHandler(db).Handle)
		prisms.POST("", di.InitializeCreatePrismHandler(db).Handle)
		prisms.POST("/:prismId/activate", di.InitializeActivatePrismHandler(db).Handle)
		prisms.PATCH("/:prismId", di.InitializeUpdatePrismHandler(db).Handle)
		prisms.DELETE("/:prismId", di.InitializeDeletePrismHandler(db).Handle)
	}
}
