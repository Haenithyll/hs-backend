package response

import (
	"hs-backend/internal/model/enum"
	"time"

	"github.com/google/uuid"
)

type RefractedFacetResponseEnrichedConfigItemCommunicationService struct {
	ID      uint8                     `json:"id"`
	Name    string                    `json:"name"`
	Value   string                    `json:"value"`
	Service enum.CommunicationService `json:"service"`
}

type RefractedFacetResponseEnrichedConfigItem struct {
	CommunicationService RefractedFacetResponseEnrichedConfigItemCommunicationService `json:"communicationService"`
	Status               enum.FacetStatus                                             `json:"status"`
}

type RefractedFacetResponseEnrichedConfig struct {
	Items []RefractedFacetResponseEnrichedConfigItem `json:"items"`
}

type RefractedFacet struct {
	Id            uint8                                 `json:"id"`
	Label         string                                `json:"label"`
	Color         string                                `json:"color"`
	Configuration *RefractedFacetResponseEnrichedConfig `json:"configuration"`
	LastUpdatedAt time.Time                             `json:"lastUpdatedAt"`
}

type RefractedFacetResponseUser struct {
	ID        uuid.UUID `json:"id"`
	AvatarUrl *string   `json:"avatarUrl"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

type RefractedFacetResponse struct {
	User           RefractedFacetResponseUser `json:"user"`
	RefractedFacet *RefractedFacet            `json:"refractedFacet"`
}

type RefractedFacetResponses []RefractedFacetResponse
