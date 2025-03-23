package routes

import (
	"hs-backend/internal/handler"
	facetHandler "hs-backend/internal/handler/facet"
	userHandler "hs-backend/internal/handler/user"
	"hs-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	authenticated := api.Group("/")
	authenticated.Use(middleware.AuthMiddleware())

	deps := &handler.HandlerDeps{DB: db}

	users := authenticated.Group("/users")
	{
		users.GET("", userHandler.NewGetUserByEmailHandler(deps).Handle)
	}

	facets := authenticated.Group("/facets")
	{
		facets.GET("", facetHandler.NewGetFacetsHandler(deps).Handle)
		facets.POST("", facetHandler.NewCreateFacetHandler(deps).Handle)
	}
}
