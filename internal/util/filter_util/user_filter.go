package filter_util

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
)

func FilterUsers(users []model.User, userId uuid.UUID) []model.User {
	filteredUsers := make([]model.User, 0, len(users)-1)
	for _, user := range users {
		if user.ID == userId {
			continue
		}
		filteredUsers = append(filteredUsers, user)
	}

	return filteredUsers
}
