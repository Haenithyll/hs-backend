package query

import (
	"hs-backend/internal/user"
	"hs-backend/internal/user/dto"
)

type GetUserByEmailHandler struct {
	Repo user.Repository
}

func NewGetUserByEmailHandler(repo user.Repository) *GetUserByEmailHandler {
	return &GetUserByEmailHandler{repo}
}

func (h *GetUserByEmailHandler) Handle(input dto.GetUserByEmailInput) (*dto.GetUserByEmailResponse, error) {
	u, err := h.Repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	return &dto.GetUserByEmailResponse{
		ID:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		AvatarURL: func() string {
			if u.AvatarURL != nil {
				return *u.AvatarURL
			}
			return ""
		}(),
	}, nil
}
