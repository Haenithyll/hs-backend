package dto

type DeleteUserCommunicationServiceInput struct {
	UserCommunicationServiceID uint8 `uri:"userCommunicationServiceId" binding:"required"`
}
