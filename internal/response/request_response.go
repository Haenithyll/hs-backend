package response

import "time"

type RequestResponse struct {
	ID        uint8                 `json:"id"`
	Issuer    *UserResponse         `json:"issuer"`
	Receiver  *UserResponse         `json:"receiver"`
	Topic     string                `json:"topic"`
	Level     *RequestLevelResponse `json:"level"`
	IsRead    bool                  `json:"isRead"`
	ReadAt    *time.Time            `json:"readAt,omitempty"`
	CreatedAt time.Time             `json:"createdAt"`
}

type RequestResponses []RequestResponse
