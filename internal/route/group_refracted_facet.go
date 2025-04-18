package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRefractedFacetRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	// refractedFacetHandler := di.InitializeRefractedFacetHandler(db)

	// refractedFacets := rg.Group("/refracted-facets")
	// {
	// 	refractedFacets.GET("", refractedFacetHandler.GetRefractedFacets)
	// }
}
