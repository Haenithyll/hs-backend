package dto

import (
	"errors"

	"hs-backend/internal/model/json"
	"hs-backend/internal/util"

	"github.com/google/uuid"
)

type CreatePrismInput struct {
	Name          string           `json:"name" binding:"required"`
	Configuration json.PrismConfig `json:"configuration" binding:"required"`
}

func (i *CreatePrismInput) Validate() error {
	userIds := make([]uuid.UUID, len(i.Configuration.Users))

	if len(i.Configuration.Users) == 0 {
		return nil
	}

	for index, user := range i.Configuration.Users {
		userIds[index] = user.UserId
	}

	if !util.AreUUIDsUnique(userIds) {
		return errors.New("user IDs must be unique")
	}

	return nil
}
