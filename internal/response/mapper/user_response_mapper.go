package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/response"
)

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
