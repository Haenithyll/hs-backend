package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/response"
)

func ToRequestResponsesFromEnrichedRequests(requests []model.Request) response.RequestResponses {
	responses := make([]response.RequestResponse, len(requests))
	for i, request := range requests {
		responses[i] = *ToRequestResponseFromEnrichedRequest(&request)
	}

	return responses
}

func ToRequestResponseFromEnrichedRequest(request *model.Request) *response.RequestResponse {
	return &response.RequestResponse{
		ID:        request.ID,
		Issuer:    ToUserResponse(request.Issuer),
		Receiver:  ToUserResponse(request.Receiver),
		Topic:     request.Topic,
		Level:     ToRequestLevelResponse(request.Level),
		IsRead:    request.IsRead,
		ReadAt:    request.ReadAt,
		CreatedAt: request.CreatedAt,
	}
}
