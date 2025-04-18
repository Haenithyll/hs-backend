package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserCommunicationServiceRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userCommunicationServiceHandler := di.InitializeUserCommunicationServiceHandler(db)

	userCommunicationServices := rg.Group("/users/communication-services")
	{
		userCommunicationServices.GET("", userCommunicationServiceHandler.GetAll)
		userCommunicationServices.POST("", userCommunicationServiceHandler.Create)
		userCommunicationServices.PATCH(":userCommunicationServiceId", userCommunicationServiceHandler.Update)
		userCommunicationServices.DELETE(":userCommunicationServiceId", userCommunicationServiceHandler.Delete)
	}
}
