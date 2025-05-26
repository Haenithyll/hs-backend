package response

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarURL string `json:"avatarUrl,omitempty"`
	Email     string `json:"email"`
}

type UserResponses []UserResponse