package enum

import (
	"database/sql/driver"
	"fmt"
)

type FacetStatus string

const (
	Available     FacetStatus = "available"
	EmergencyOnly FacetStatus = "emergencyOnly"
)

func (fs FacetStatus) IsValid() bool {
	switch fs {
	case Available, EmergencyOnly:
		return true
	default:
		return false
	}
}

func (fs *FacetStatus) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		*fs = FacetStatus(string(v))
	case string:
		*fs = FacetStatus(v)
	default:
		return fmt.Errorf("failed to scan FacetStatus: %v", value)
	}
	return nil
}

func (fs FacetStatus) Value() (driver.Value, error) {
	return string(fs), nil
}
