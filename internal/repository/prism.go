package repository

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrismRepository interface {
	FindManyByUserId(userId uuid.UUID) ([]model.Prism, error)
	FindOneByIDAndUserID(id uint8, userId uuid.UUID) (*model.Prism, error)

	CreateOne(prism *model.Prism) error
	UpdateOne(prism *model.Prism) error

	DeleteOneByIDAndUserID(id uint8, userId uuid.UUID) error
}

type prismRepository struct {
	db *gorm.DB
}

func NewPrismRepository(db *gorm.DB) PrismRepository {
	return &prismRepository{db}
}

func (r *prismRepository) FindManyByUserId(userId uuid.UUID) ([]model.Prism, error) {
	var prismList []model.Prism
	if err := r.db.
		Model(&model.Prism{}).
		Where("user_id = ?", userId).
		Find(&prismList).Error; err != nil {
		return nil, err
	}
	return prismList, nil
}

func (r *prismRepository) FindOneByIDAndUserID(id uint8, userId uuid.UUID) (*model.Prism, error) {
	var prism model.Prism
	if err := r.db.First(&prism, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		return nil, err
	}
	return &prism, nil
}

func (r *prismRepository) CreateOne(prism *model.Prism) error {
	if err := r.db.Create(prism).Error; err != nil {
		return err
	}
	return nil
}

func (r *prismRepository) UpdateOne(prism *model.Prism) error {
	return r.db.
		Model(&model.Prism{}).
		Where("id = ? AND user_id = ?", prism.ID, prism.UserId).
		Updates(map[string]interface{}{
			"name":          prism.Name,
			"configuration": prism.Configuration,
		}).Error
}

func (r *prismRepository) DeleteOneByIDAndUserID(id uint8, userId uuid.UUID) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&model.Prism{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
