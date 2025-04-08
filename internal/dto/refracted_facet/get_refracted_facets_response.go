package dto

import (
	"hs-backend/internal/model"
	"hs-backend/internal/types"

	"github.com/google/uuid"
)

type GetRefractedFacetsResponse []types.RefractedUserFacet

func ToGetRefractedFacetsResponse(
	refractedFacetByUserIdMap map[uuid.UUID]types.RefractedFacet,
	users []model.User,
	userId uuid.UUID,
) *GetRefractedFacetsResponse {
	response := make(GetRefractedFacetsResponse, 0, len(users))

	for _, user := range users {
		if user.ID == userId {
			continue
		}

		var refractedFacet *types.RefractedFacet = nil

		if rf, ok := refractedFacetByUserIdMap[user.ID]; ok {
			refractedFacet = &rf
		}

		response = append(response, types.RefractedUserFacet{
			User:           &user,
			RefractedFacet: refractedFacet,
		})
	}

	return &response
}
