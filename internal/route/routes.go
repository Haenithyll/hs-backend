package route

import (
	"hs-backend/internal/handler"
	facetHandler "hs-backend/internal/handler/facet"
	prismHandler "hs-backend/internal/handler/prism"
	userHandler "hs-backend/internal/handler/user"
	userCommunicationServiceHandler "hs-backend/internal/handler/user_communication_service"
	"hs-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	authenticated := api.Group("/")
	authenticated.Use(middleware.AuthMiddleware())

	deps := &handler.HandlerDeps{DB: db}

	facets := authenticated.Group("/facets")
	{
		facets.GET("", facetHandler.NewGetFacetsHandler(deps).Handle)
		facets.POST("", facetHandler.NewCreateFacetHandler(deps).Handle)
		facets.PATCH(":facetId", facetHandler.NewUpdateFacetHandler(deps).Handle)
		facets.DELETE(":facetId", facetHandler.NewDeleteFacetHandler(deps).Handle)
	}

	prisms := authenticated.Group("/prisms")
	{
		prisms.POST("", prismHandler.NewCreatePrismHandler(deps).Handle)
	}

	users := authenticated.Group("/users")
	{
		users.GET("", userHandler.NewGetUserByEmailHandler(deps).Handle)
		users.GET("/me", userHandler.NewGetUserMeHandler(deps).Handle)
	}

	userCommunicationServices := authenticated.Group("/users/communication-services")
	{
		userCommunicationServices.GET("", userCommunicationServiceHandler.NewGetUserCommunicationServiceHandler(deps).Handle)
		userCommunicationServices.POST("", userCommunicationServiceHandler.NewCreateUserCommunicationServiceHandler(deps).Handle)
		userCommunicationServices.PATCH(":userCommunicationServiceId", userCommunicationServiceHandler.NewUpdateUserCommunicationServiceHandler(deps).Handle)
		userCommunicationServices.DELETE(":userCommunicationServiceId", userCommunicationServiceHandler.NewDeleteUserCommunicationServiceHandler(deps).Handle)
	}
}
