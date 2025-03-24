package model

import (
	"time"

	"hs-backend/internal/model/json"

	"github.com/google/uuid"
)

type Facet struct {
	ID            int8             `gorm:"primaryKey;column=id"`
	Color         string           `gorm:"column=color"`
	UserId        uuid.UUID        `gorm:"type:uuid;column=user_id"`
	User          *User            `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	PublicLabel   string           `gorm:"column=public_label"`
	PrivateLabel  string           `gorm:"column=private_label"`
	Configuration json.FacetConfig `gorm:"type:jsonb;column=configuration"`
	CreatedAt     time.Time        `gorm:"column=created_at"`
}

func (Facet) TableName() string {
	return "facets"
}
