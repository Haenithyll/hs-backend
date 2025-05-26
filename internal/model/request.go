package model

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID         uint8         `gorm:"primaryKey;column=id"`
	IssuerId   uuid.UUID     `gorm:"type:uuid;column=issuer_id"`
	Issuer     *User         `gorm:"foreignKey:IssuerId;constraint:OnDelete:CASCADE"`
	ReceiverId uuid.UUID     `gorm:"type:uuid;column=receiver_id"`
	Receiver   *User         `gorm:"foreignKey:ReceiverId;constraint:OnDelete:CASCADE"`
	Topic      string        `gorm:"type:string;column=topic"`
	LevelId    uint8         `gorm:"type:uint8;column=level_id"`
	Level      *RequestLevel `gorm:"foreignKey:LevelId;constraint:OnDelete:CASCADE"`
	IsRead     bool          `gorm:"type:boolean;column=is_read"`
	ReadAt     *time.Time    `gorm:"type:timestamp;column=read_at"`
	CreatedAt  time.Time     `gorm:"type:timestamp;column=created_at"`
}

func (Request) TableName() string {
	return "requests"
}
