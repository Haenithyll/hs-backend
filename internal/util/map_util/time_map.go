package map_util

import (
	"hs-backend/internal/model"
	"time"

	"github.com/google/uuid"
)

func BuildLastUpdatedAtMapByUserId(userPrismTrackers []model.UserPrismTracker) map[uuid.UUID]time.Time {
	m := make(map[uuid.UUID]time.Time)

	for _, userPrismTracker := range userPrismTrackers {
		m[userPrismTracker.UserId] = userPrismTracker.LastUpdatedAt
	}

	return m
}
