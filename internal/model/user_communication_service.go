package model

import (
	"hs-backend/internal/model/enum"
	"time"

	"github.com/google/uuid"
)

type UserCommunicationService struct {
	ID        uint8                     `gorm:"primaryKey;column=id"`
	UserId    uuid.UUID                 `gorm:"type:uuid;column=user_id"`
	User      *User                     `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Name      string                    `gorm:"column=name"`
	Value     string                    `gorm:"column=value"`
	Service   enum.CommunicationService `gorm:"type:communication_service;column=service"`
	CreatedAt time.Time                 `gorm:"column=created_at"`
}

func (UserCommunicationService) TableName() string {
	return "user_communication_services"
}
