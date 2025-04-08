package types

import (
	"hs-backend/internal/model"
)

type RefractedUserFacet struct {
	User           *model.User     `json:"user"`
	RefractedFacet *RefractedFacet `json:"refractedFacet"`
}
