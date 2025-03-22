package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;column=id"`
	FirstName      string         `gorm:"column=first_name"`
	LastName       string         `gorm:"column=last_name"`
	AvatarURL      *string        `gorm:"column=avatar_url"` // nullable
	BillingAddress datatypes.JSON `gorm:"column=billing_address"`
	PaymentMethod  datatypes.JSON `gorm:"column=payment_method"`
	Email          string         `gorm:"column=email"`
}

func (User) TableName() string {
	return "users"
}
