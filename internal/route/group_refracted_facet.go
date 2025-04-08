package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRefractedFacetRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	refractedFacets := rg.Group("/refracted-facets")
	{
		refractedFacets.GET("", di.InitializeGetRefractedFacetsHandler(db).Handle)
	}
}
