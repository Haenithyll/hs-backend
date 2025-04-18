package map_util

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
)

func BuildUserMapById(users []model.User) map[uuid.UUID]model.User {
	m := make(map[uuid.UUID]model.User, len(users))
	for _, user := range users {
		m[user.ID] = user
	}
	return m
}
