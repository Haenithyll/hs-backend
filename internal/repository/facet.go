package repository

import (
	dto "hs-backend/internal/dto/facet"
	"hs-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FacetRepository interface {
	FindByUserId(userId uuid.UUID) ([]model.Facet, error)
	Create(input dto.CreateFacetInput, userId uuid.UUID) (*model.Facet, error)
}

type facetRepository struct {
	db *gorm.DB
}

func NewFacetRepository(db *gorm.DB) FacetRepository {
	return &facetRepository{db}
}

func (r *facetRepository) FindByUserId(userId uuid.UUID) ([]model.Facet, error) {
	var facets []model.Facet
	if err := r.db.Where("user_id = ?", userId).Find(&facets).Error; err != nil {
		return nil, err
	}
	return facets, nil
}

func (r *facetRepository) Create(input dto.CreateFacetInput, userId uuid.UUID) (*model.Facet, error) {
	facet := model.Facet{
		Color:         input.Color,
		PublicLabel:   input.PublicLabel,
		PrivateLabel:  input.PrivateLabel,
		Configuration: input.Configuration,
		UserId:        userId,
	}

	if err := r.db.Create(&facet).Error; err != nil {
		return nil, err
	}

	return &facet, nil
}
