package enum

import (
	"database/sql/driver"
	"fmt"
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

func (cs CommunicationService) IsValid() bool {
	switch cs {
	case FaceToFace, Phone, Message, Email, Discord, MicrosoftTeams:
		return true
	default:
		return false
	}
}

func (cs *CommunicationService) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		*cs = CommunicationService(string(v))
	case string:
		*cs = CommunicationService(v)
	default:
		return fmt.Errorf("failed to scan CommunicationService: %v", value)
	}
	return nil
}

func (cs CommunicationService) Value() (driver.Value, error) {
	return string(cs), nil
}
