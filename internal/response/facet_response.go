package response

import (
	"hs-backend/internal/model/enum"
	"time"
)

type FacetResponseEnrichedConfigItemCommunicationService struct {
	ID      uint8                     `json:"id"`
	Name    string                    `json:"name"`
	Value   string                    `json:"value"`
	Service enum.CommunicationService `json:"service"`
}

type FacetResponseEnrichedConfigItem struct {
	CommunicationService FacetResponseEnrichedConfigItemCommunicationService `json:"communicationService"`
	Status               enum.FacetStatus                                    `json:"status"`
}

type FacetResponseEnrichedConfig struct {
	Items []FacetResponseEnrichedConfigItem `json:"items"`
}

type FacetResponse struct {
	ID            uint8                       `json:"id"`
	Color         string                      `json:"color"`
	PublicLabel   string                      `json:"publicLabel"`
	PrivateLabel  string                      `json:"privateLabel"`
	Configuration FacetResponseEnrichedConfig `json:"configuration"`
	CreatedAt     time.Time                   `json:"createdAt"`
}

type FacetResponses []FacetResponse
