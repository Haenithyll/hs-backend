package service

import (
	"hs-backend/internal/repository"
)

type RefractedFacetService struct {
	FacetRepository            repository.FacetRepository
	UserRepository             repository.UserRepository
	UserPrismTrackerRepository repository.UserPrismTrackerRepository
}

func NewRefractedFacetService(
	facetRepository repository.FacetRepository,
	userRepository repository.UserRepository,
	userPrismTrackerRepository repository.UserPrismTrackerRepository,
) *RefractedFacetService {
	return &RefractedFacetService{
		facetRepository,
		userRepository,
		userPrismTrackerRepository,
	}
}
