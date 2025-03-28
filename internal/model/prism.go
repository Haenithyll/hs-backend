package model

import (
	"time"

	"hs-backend/internal/model/json"

	"github.com/google/uuid"
)

type Prism struct {
	ID            uint8            `gorm:"primaryKey;column=id"`
	Name          string           `gorm:"column=name"`
	UserId        uuid.UUID        `gorm:"type:uuid;column=user_id"`
	User          *User            `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Configuration json.PrismConfig `gorm:"type:jsonb;column=configuration"`
	CreatedAt     time.Time        `gorm:"column=created_at"`
}

func (Prism) TableName() string {
	return "prisms"
}
