package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type FacetStatus string

const (
	Available     FacetStatus = "available"
	EmergencyOnly FacetStatus = "emergency-only"
)

func (fs *FacetStatus) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan FacetStatus: %v", value)
	}
	*fs = FacetStatus(string(bytes))
	return nil
}

func (fs FacetStatus) Value() (driver.Value, error) {
	return string(fs), nil
}

type FacetConfigItem struct {
	Id     uint8       `json:"id"`
	Status FacetStatus `json:"status"`
}

type FacetConfig struct {
	Items []FacetConfigItem `json:"items"`
}

func (fc *FacetConfig) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan FacetConfig: %v", value)
	}
	return json.Unmarshal(bytes, fc) // note: unmarshalling into `fc` itself
}

func (fc FacetConfig) Value() (driver.Value, error) {
	return json.Marshal(fc)
}

type Facet struct {
	ID            int8        `gorm:"primaryKey;column=id" json:"id"`
	Color         string      `gorm:"column=color" json:"color"`
	UserId        uuid.UUID   `gorm:"type:uuid;column=user_id" json:"userId"`
	User          *User       `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	PublicLabel   string      `gorm:"column=public_label" json:"publicLabel"`
	PrivateLabel  string      `gorm:"column=private_label" json:"privateLabel"`
	Configuration FacetConfig `gorm:"type:jsonb;column=configuration" json:"configuration"`
	CreatedAt     time.Time   `gorm:"column=created_at" json:"createdAt"`
}

func (Facet) TableName() string {
	return "facets"
}
