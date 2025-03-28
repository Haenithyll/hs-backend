package repository

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCommunicationServiceRepository interface {
	FindManyByUserId(userId uuid.UUID) ([]model.UserCommunicationService, error)
	FindOneByIDAndUserID(id uint8, userId uuid.UUID) (*model.UserCommunicationService, error)

	CreateOne(ucs *model.UserCommunicationService) error
	UpdateOne(ucs *model.UserCommunicationService) error

	DeleteOneByIDAndUserID(id uint8, userId uuid.UUID) error
}

type userCommunicationServiceRepository struct {
	db *gorm.DB
}

func NewUserCommunicationServiceRepository(db *gorm.DB) UserCommunicationServiceRepository {
	return &userCommunicationServiceRepository{db}
}

func (r *userCommunicationServiceRepository) FindManyByUserId(userId uuid.UUID) ([]model.UserCommunicationService, error) {
	var ucsList []model.UserCommunicationService
	if err := r.db.
		Model(&model.UserCommunicationService{}).
		Where("user_id = ?", userId).
		Find(&ucsList).Error; err != nil {
		return nil, err
	}
	return ucsList, nil
}

func (r *userCommunicationServiceRepository) FindOneByIDAndUserID(id uint8, userId uuid.UUID) (*model.UserCommunicationService, error) {
	var ucs model.UserCommunicationService
	if err := r.db.First(&ucs, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		return nil, err
	}
	return &ucs, nil
}

func (r *userCommunicationServiceRepository) CreateOne(ucs *model.UserCommunicationService) error {
	if err := r.db.Create(ucs).Error; err != nil {
		return err
	}
	return nil
}

func (r *userCommunicationServiceRepository) UpdateOne(ucs *model.UserCommunicationService) error {
	return r.db.
		Model(&model.UserCommunicationService{}).
		Where("id = ? AND user_id = ?", ucs.ID, ucs.UserId).
		Updates(map[string]interface{}{
			"name":    ucs.Name,
			"value":   ucs.Value,
			"service": ucs.Service,
		}).Error
}

func (r *userCommunicationServiceRepository) DeleteOneByIDAndUserID(id uint8, userId uuid.UUID) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&model.UserCommunicationService{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
