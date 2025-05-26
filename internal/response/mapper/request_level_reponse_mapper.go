package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/response"
)

func ToRequestLevelResponses(levels []model.RequestLevel) response.RequestLevelResponses {
	responses := make(response.RequestLevelResponses, len(levels))
	for i, level := range levels {
		responses[i] = *ToRequestLevelResponse(&level)
	}
	return responses
}

func ToRequestLevelResponse(level *model.RequestLevel) *response.RequestLevelResponse {
	return &response.RequestLevelResponse{
		ID:    level.ID,
		Label: level.Label,
	}
}
