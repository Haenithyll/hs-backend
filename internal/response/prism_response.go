package response

import (
	"time"

	"github.com/google/uuid"
)

type PrismResponseFacet struct {
	ID           uint8  `json:"id"`
	Color        string `json:"color"`
	PublicLabel  string `json:"publicLabel"`
	PrivateLabel string `json:"privateLabel"`
}

type PrismResponseUser struct {
	ID        uuid.UUID `json:"id"`
	AvatarUrl *string   `json:"avatarUrl"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

type PrismResponseEnrichedUserItem struct {
	User  PrismResponseUser  `json:"user"`
	Facet PrismResponseFacet `json:"facet"`
}

type PrismResponseEnrichedConfig struct {
	Base  PrismResponseFacet              `json:"base"`
	Users []PrismResponseEnrichedUserItem `json:"users"`
}

type PrismResponse struct {
	ID            uint8                       `json:"id"`
	Name          string                      `json:"name"`
	Configuration PrismResponseEnrichedConfig `json:"configuration"`
	CreatedAt     time.Time                   `json:"createdAt"`
}

type PrismResponses []PrismResponse
