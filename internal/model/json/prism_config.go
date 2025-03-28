package json

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type PrismConfigUserItem struct {
	UserId  uuid.UUID `json:"userId"`
	FacetId uint8     `json:"facetId"`
}

type PrismConfig struct {
	Base  uint8                 `json:"base"`
	Users []PrismConfigUserItem `json:"users"`
}

func (pc *PrismConfig) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan PrismConfig: %v", value)
	}
	return json.Unmarshal(bytes, pc)
}

func (pc PrismConfig) Value() (driver.Value, error) {
	return json.Marshal(pc)
}
