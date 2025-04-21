package map_util

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
)

func BuildUserCommunicationServiceMapById(userCommunicationServices []model.UserCommunicationService) map[uint8]model.UserCommunicationService {
	m := make(map[uint8]model.UserCommunicationService, len(userCommunicationServices))
	for _, userCommunicationService := range userCommunicationServices {
		m[userCommunicationService.ID] = userCommunicationService
	}
	return m
}

func BuildUserCommunicationServiceMapByUserId(
	userCommunicationServices []model.UserCommunicationService,
) map[uuid.UUID][]model.UserCommunicationService {
	m := make(map[uuid.UUID][]model.UserCommunicationService)
	for _, userCommunicationService := range userCommunicationServices {
		m[userCommunicationService.UserId] =
			append(m[userCommunicationService.UserId], userCommunicationService)
	}
	return m
}
