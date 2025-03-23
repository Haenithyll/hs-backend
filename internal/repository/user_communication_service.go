package repository

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCommunicationServiceRepository interface {
	FindIDsByUserId(userId uuid.UUID) ([]uint8, error)
}

type userCommunicationServiceRepository struct {
	db *gorm.DB
}

func NewUserCommunicationServiceRepository(db *gorm.DB) UserCommunicationServiceRepository {
	return &userCommunicationServiceRepository{db}
}

func (r *userCommunicationServiceRepository) FindIDsByUserId(userId uuid.UUID) ([]uint8, error) {
	var ids []uint8
	if err := r.db.
		Model(&model.UserCommunicationService{}).
		Where("user_id = ?", userId).
		Pluck("id", &ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
