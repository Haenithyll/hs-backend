package dto

import "hs-backend/internal/model"

type GetUserMeResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarURL string `json:"avatarUrl,omitempty"`
	Email     string `json:"email"`
}

func ToGetUserMeResponse(user *model.User) *GetUserMeResponse {
	return &GetUserMeResponse{
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
