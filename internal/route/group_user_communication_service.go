package route

import (
	di "hs-backend/internal/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserCommunicationServiceRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userCommunicationServices := rg.Group("/users/communication-services")
	{
		userCommunicationServices.GET("", di.InitializeGetUserCommunicationServiceHandler(db).Handle)
		userCommunicationServices.POST("", di.InitializeCreateUserCommunicationServiceHandler(db).Handle)
		userCommunicationServices.PATCH(":userCommunicationServiceId", di.InitializeUpdateUserCommunicationServiceHandler(db).Handle)
		userCommunicationServices.DELETE(":userCommunicationServiceId", di.InitializeDeleteUserCommunicationServiceHandler(db).Handle)
	}
}
