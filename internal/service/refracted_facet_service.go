package service

import (
	"hs-backend/internal/domain"
	"hs-backend/internal/model"
	"hs-backend/internal/repository"
	"hs-backend/internal/response"
	"hs-backend/internal/response/mapper"
	"hs-backend/internal/util/filter_util"
	"hs-backend/internal/util/map_util"

	"github.com/google/uuid"
)

type RefractedFacetService struct {
	FacetRepository                    repository.FacetRepository
	UserCommunicationServiceRepository repository.UserCommunicationServiceRepository
	UserRepository                     repository.UserRepository
	UserPrismTrackerRepository         repository.UserPrismTrackerRepository
}

func NewRefractedFacetService(
	facetRepository repository.FacetRepository,
	userCommunicationServiceRepository repository.UserCommunicationServiceRepository,
	userRepository repository.UserRepository,
	userPrismTrackerRepository repository.UserPrismTrackerRepository,
) *RefractedFacetService {
	return &RefractedFacetService{
		facetRepository,
		userCommunicationServiceRepository,
		userRepository,
		userPrismTrackerRepository,
	}
}

// Fetches all refracted facets for every user except the current user.
// It also retrieves related user communication services for DTO conversion.
func (s *RefractedFacetService) GetRefractedFacets(userId uuid.UUID) (response.RefractedFacetResponses, *domain.DomainError) {
	userPrismTrackers, err := s.UserPrismTrackerRepository.FindAllWithPrismsAndUsers()
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get user prism trackers: "+err.Error())
	}

	if len(userPrismTrackers) == 0 {
		return []response.RefractedFacetResponse{}, nil
	}

	users := make([]model.User, 0, len(userPrismTrackers))
	for _, userPrismTracker := range userPrismTrackers {
		users = append(users, *userPrismTracker.User)
	}

	filteredUsers := filter_util.FilterUsers(users, userId)

	lastUpdatedAtByUserIdMap := map_util.BuildLastUpdatedAtMapByUserId(userPrismTrackers)

	refractedFacetIds := s.determineRefractedFacetIds(userId, userPrismTrackers)

	facets, err := s.FacetRepository.FindManyRefractedByIds(refractedFacetIds)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get facets: "+err.Error())
	}

	facetByUserIdMap, userCommunicationServiceIds := s.mapFacetsByUserIdAndExtractUserCommunicationServiceIds(facets)

	userCommunicationServices, err := s.UserCommunicationServiceRepository.FindManyByIds(userCommunicationServiceIds)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to get user communication services: "+err.Error())
	}

	userCommunicationServicesByUserIdMap := map_util.BuildUserCommunicationServiceMapByUserId(userCommunicationServices)

	return mapper.ToRefractedFacetResponses(
		filteredUsers,
		facetByUserIdMap,
		userCommunicationServicesByUserIdMap,
		lastUpdatedAtByUserIdMap,
	), nil
}

// Determines the refracted facet IDs of every user except the current user.
func (s *RefractedFacetService) determineRefractedFacetIds(userId uuid.UUID, userPrismTrackers []model.UserPrismTracker) []uint8 {
	refractedFacetIds := make([]uint8, 0, len(userPrismTrackers))

	determineRefractedFacetId := func(userPrismTracker model.UserPrismTracker) uint8 {
		refractedFacetId := userPrismTracker.Prism.Configuration.Base

		if len(userPrismTracker.Prism.Configuration.Users) > 0 {
			for _, configUserItem := range userPrismTracker.Prism.Configuration.Users {
				if configUserItem.UserId == userId {
					refractedFacetId = configUserItem.FacetId
					break
				}
			}

		}

		return refractedFacetId
	}

	for _, userPrismTracker := range userPrismTrackers {
		if userPrismTracker.UserId == userId {
			continue
		}

		refractedFacetIds = append(refractedFacetIds, determineRefractedFacetId(userPrismTracker))
	}

	return refractedFacetIds
}

// Maps the facets by user ID and extracts the user communication service IDs from the facet configuration.
func (s *RefractedFacetService) mapFacetsByUserIdAndExtractUserCommunicationServiceIds(
	facets []model.Facet,
) (map[uuid.UUID]model.Facet, []uint8) {
	facetByUserIdMap := make(map[uuid.UUID]model.Facet)
	userCommunicationServiceIds := make([]uint8, 0)

	for _, facet := range facets {
		facetByUserIdMap[facet.UserId] = facet

		for _, item := range facet.Configuration.Items {
			userCommunicationServiceIds = append(userCommunicationServiceIds, item.Id)
		}
	}

	return facetByUserIdMap, userCommunicationServiceIds
}
