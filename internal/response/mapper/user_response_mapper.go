package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/response"
)

func ToUserResponses(users []model.User) response.UserResponses {
	responses := make([]response.UserResponse, len(users))

	for i, user := range users {
		responses[i] = *ToUserResponse(&user)
	}

	return responses
}

func ToUserResponse(user *model.User) *response.UserResponse {
	return &response.UserResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		AvatarURL: func() string {
			if user.AvatarURL != nil {
				return *user.AvatarURL
			}
			return ""
		}(),
		Email: user.Email,
	}
}
