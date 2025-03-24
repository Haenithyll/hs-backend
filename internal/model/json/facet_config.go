package json

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"hs-backend/internal/model/enum"
)

type FacetConfigItem struct {
	Id     uint8            `json:"id"`
	Status enum.FacetStatus `json:"status"`
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
