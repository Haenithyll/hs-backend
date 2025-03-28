package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPrismRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	prisms := rg.Group("/prisms")
	{
		prisms.POST("", di.InitializeCreatePrismHandler(db).Handle)
	}
}
