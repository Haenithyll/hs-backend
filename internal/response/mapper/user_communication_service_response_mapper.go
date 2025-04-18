package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/response"
)

func ToUserCommunicationServiceResponses(userCommunicationServices []model.UserCommunicationService) response.UserCommunicationServiceResponses {
	responses := make([]response.UserCommunicationServiceResponse, len(userCommunicationServices))
	for i, userCommunicationService := range userCommunicationServices {
		responses[i] = ToUserCommunicationServiceResponse(userCommunicationService)
	}
	return responses
}

func ToUserCommunicationServiceResponse(userCommunicationService model.UserCommunicationService) response.UserCommunicationServiceResponse {
	return response.UserCommunicationServiceResponse{
		ID:        userCommunicationService.ID,
		Name:      userCommunicationService.Name,
		Value:     userCommunicationService.Value,
		Service:   userCommunicationService.Service,
		CreatedAt: userCommunicationService.CreatedAt,
	}
}
