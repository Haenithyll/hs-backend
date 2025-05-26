package mapper

import (
	"hs-backend/internal/response"
)

func ToKeyResponse(key string) *response.KeyResponse {
	return &response.KeyResponse{Key: key}
}
