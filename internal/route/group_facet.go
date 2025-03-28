package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterFacetRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	facets := rg.Group("/facets")
	{
		facets.GET("", di.InitializeGetFacetsHandler(db).Handle)
		facets.POST("", di.InitializeCreateFacetHandler(db).Handle)
		facets.PATCH(":facetId", di.InitializeUpdateFacetHandler(db).Handle)
		facets.DELETE(":facetId", di.InitializeDeleteFacetHandler(db).Handle)
	}
}
