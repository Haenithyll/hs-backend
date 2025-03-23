package model

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type CommunicationService string

const (
	FaceToFace     CommunicationService = "face-to-face"
	Phone          CommunicationService = "phone"
	Message        CommunicationService = "message"
	Email          CommunicationService = "email"
	Discord        CommunicationService = "discord"
	MicrosoftTeams CommunicationService = "microsoft-teams"
)

func (cs *CommunicationService) Scan(value any) error {
	*cs = CommunicationService(value.([]byte))
	return nil
}

func (cs CommunicationService) Value() (driver.Value, error) {
	return string(cs), nil
}

type UserCommunicationService struct {
	ID          uint8                `gorm:"primaryKey;column=id"`
	UserId      uuid.UUID            `gorm:"type:uuid;column=user_id"`
	User        User                 `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Name        string               `gorm:"column=name"`
	Value       string               `gorm:"column=value"`
	ServiceType CommunicationService `gorm:"type:communication_service;column=service_type"`
}

func (UserCommunicationService) TableName() string {
	return "user_communication_services"
}
