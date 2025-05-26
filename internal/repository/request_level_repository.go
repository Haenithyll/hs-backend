package repository

import (
	"hs-backend/internal/model"

	"gorm.io/gorm"
)

type RequestLevelRepository interface {
	ExistsById(id uint8) (bool, error)

	FindAll() ([]model.RequestLevel, error)
}

type requestLevelRepository struct {
	db *gorm.DB
}

func NewRequestLevelRepository(db *gorm.DB) RequestLevelRepository {
	return &requestLevelRepository{db}
}

func (r *requestLevelRepository) ExistsById(id uint8) (bool, error) {
	var count int64
	if err := r.db.Model(&model.RequestLevel{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *requestLevelRepository) FindAll() ([]model.RequestLevel, error) {
	var requestLevels []model.RequestLevel
	if err := r.db.Find(&requestLevels).Error; err != nil {
		return nil, err
	}
	return requestLevels, nil
}
