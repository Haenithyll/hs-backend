package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterFacetRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	facetHandler := di.InitializeFacetHandler(db)

	facets := rg.Group("/facets")
	{
		facets.GET("", facetHandler.GetAll)
		facets.POST("", facetHandler.Create)
		facets.PATCH(":facetId", facetHandler.Update)
		facets.DELETE(":facetId", facetHandler.Delete)
	}
}
