package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRequestRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	requestHandler := di.InitializeRequestHandler(db)

	requests := rg.Group("/requests")
	{
		requests.GET("/received", requestHandler.GetAllReceived)
		requests.GET("/issued", requestHandler.GetAllIssued)
		requests.POST("", requestHandler.Create)
		requests.PUT("/:requestId/read", requestHandler.MarkAsRead)
	}
}
