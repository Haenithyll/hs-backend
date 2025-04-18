package response

import (
	"hs-backend/internal/model/enum"
	"time"
)

type UserCommunicationServiceResponse struct {
	ID        uint8                     `json:"id"`
	Name      string                    `json:"name"`
	Value     string                    `json:"value"`
	Service   enum.CommunicationService `json:"service"`
	CreatedAt time.Time                 `json:"createdAt"`
}

type UserCommunicationServiceResponses []UserCommunicationServiceResponse
