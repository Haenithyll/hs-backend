package mapper

import (
	"hs-backend/internal/model"
	"hs-backend/internal/response"
	"hs-backend/internal/types"

	"github.com/google/uuid"
)

func ToRefractedFacetResponse(
	refractedFacetByUserIdMap map[uuid.UUID]types.RefractedFacet,
	users []model.User,
	userId uuid.UUID,
) *response.RefractedFacetResponse {
	response := make(response.RefractedFacetResponse, 0, len(users))

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
