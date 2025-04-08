package dto

import (
	"errors"

	"hs-backend/internal/model/json"
	"hs-backend/internal/util"

	"github.com/google/uuid"
)

type UpdatePrismInput struct {
	PrismID       uint8             `uri:"prismId" json:"-" binding:"required"`
	Name          *string           `json:"name,omitempty"`
	Configuration *json.PrismConfig `json:"configuration,omitempty"`
}

func (i *UpdatePrismInput) Validate() error {
	if i.Configuration != nil {
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

	return nil
}
