package map_util

import (
	"hs-backend/internal/model"
)

func BuildUserCommunicationServiceMapById(userCommunicationServices []model.UserCommunicationService) map[uint8]model.UserCommunicationService {
	m := make(map[uint8]model.UserCommunicationService, len(userCommunicationServices))
	for _, userCommunicationService := range userCommunicationServices {
		m[userCommunicationService.ID] = userCommunicationService
	}
	return m
}
