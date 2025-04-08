package model

import (
	"time"

	"github.com/google/uuid"
)

type UserPrismTracker struct {
	ID            uint8     `gorm:"primaryKey;column=id"`
	PrismID       uint8     `gorm:"column=prism_id"`
	Prism         *Prism    `gorm:"foreignKey:PrismID;constraint:OnDelete:CASCADE"`
	UserId        uuid.UUID `gorm:"type:uuid;column=user_id"`
	User          *User     `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	LastUpdatedAt time.Time `gorm:"column=last_updated_at"`
}

func (UserPrismTracker) TableName() string {
	return "user_prism_trackers"
}
