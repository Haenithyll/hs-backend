package route

import (
	"hs-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	authenticated := api.Group("/")
	authenticated.Use(middleware.AuthMiddleware())

	RegisterKeyRoutes(authenticated, db)
	RegisterFacetRoutes(authenticated, db)
	RegisterPrismRoutes(authenticated, db)
	RegisterRefractedFacetRoutes(authenticated, db)
	RegisterRequestLevelRoutes(authenticated, db)
	RegisterRequestRoutes(authenticated, db)
	RegisterUserRoutes(authenticated, db)
	RegisterUserCommunicationServiceRoutes(authenticated, db)
}
