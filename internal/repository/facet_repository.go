package repository

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FacetRepository interface {
	FindManyByUserId(userId uuid.UUID) ([]model.Facet, error)
	FindManyRefractedByIds(ids []uint8) ([]model.Facet, error)
	FindOneByIDAndUserID(id uint8, userId uuid.UUID) (*model.Facet, error)
	FindManyByIdsAndUserId(ids []uint8, userId uuid.UUID) ([]model.Facet, error)

	CreateOne(facet *model.Facet) error
	UpdateOne(facet *model.Facet) error

	RemoveUserCommunicationServiceFromFacets(userCommunicationServiceID uint8) error

	DeleteOneByIDAndUserID(id uint8, userId uuid.UUID) error
}

type facetRepository struct {
	db *gorm.DB
}

func NewFacetRepository(db *gorm.DB) FacetRepository {
	return &facetRepository{db}
}

func (r *facetRepository) FindManyByUserId(userId uuid.UUID) ([]model.Facet, error) {
	var facets []model.Facet
	if err := r.db.Where("user_id = ?", userId).Find(&facets).Error; err != nil {
		return nil, err
	}
	return facets, nil
}

func (r *facetRepository) FindManyRefractedByIds(ids []uint8) ([]model.Facet, error) {
	if len(ids) == 0 {
		return []model.Facet{}, nil
	}

	convertedIDs := make([]uint, len(ids))
	for i, id := range ids {
		convertedIDs[i] = uint(id)
	}

	var facets []model.Facet
	if err := r.db.
		Where("id IN ?", convertedIDs).
		Select("id", "user_id", "public_label", "color", "configuration").
		Find(&facets).Error; err != nil {
		return nil, err
	}
	return facets, nil
}

func (r *facetRepository) FindOneByIDAndUserID(id uint8, userId uuid.UUID) (*model.Facet, error) {
	var facet model.Facet
	if err := r.db.First(&facet, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		return nil, err
	}
	return &facet, nil
}

func (r *facetRepository) FindManyByIdsAndUserId(ids []uint8, userId uuid.UUID) ([]model.Facet, error) {
	if len(ids) == 0 {
		return []model.Facet{}, nil
	}

	convertedIDs := make([]uint, len(ids))
	for i, id := range ids {
		convertedIDs[i] = uint(id)
	}

	var facets []model.Facet
	if err := r.db.Where("id IN ? AND user_id = ?", convertedIDs, userId).Find(&facets).Error; err != nil {
		return nil, err
	}
	return facets, nil
}

func (r *facetRepository) CreateOne(facet *model.Facet) error {
	if err := r.db.Create(&facet).Error; err != nil {
		return err
	}
	return nil
}

func (r *facetRepository) UpdateOne(facet *model.Facet) error {
	return r.db.
		Model(&model.Facet{}).
		Where("id = ? AND user_id = ?", facet.ID, facet.UserId).
		Updates(map[string]any{
			"color":         facet.Color,
			"public_label":  facet.PublicLabel,
			"private_label": facet.PrivateLabel,
			"configuration": facet.Configuration,
		}).Error
}

func (r *facetRepository) RemoveUserCommunicationServiceFromFacets(userCommunicationServiceID uint8) error {
	return r.db.Exec(`
		UPDATE public.facets
		SET configuration = jsonb_set(
			configuration,
			'{items}',
			COALESCE((
				SELECT jsonb_agg(item)
				FROM jsonb_array_elements(configuration->'items') AS item
				WHERE (item->>'id')::int != ?
			), '[]'::jsonb)
		)
		WHERE configuration->'items' IS NOT NULL;
	`, userCommunicationServiceID).Error
}

func (r *facetRepository) DeleteOneByIDAndUserID(id uint8, userId uuid.UUID) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&model.Facet{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
