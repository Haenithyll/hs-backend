package repository

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCommunicationServiceRepository interface {
	FindManyByUserId(userId uuid.UUID) ([]model.UserCommunicationService, error)
	FindManyByIdsAndUserId(ids []uint8, userId uuid.UUID) ([]model.UserCommunicationService, error)
	FindOneByIDAndUserID(id uint8, userId uuid.UUID) (*model.UserCommunicationService, error)

	CreateOne(userCommunicationService *model.UserCommunicationService) error
	UpdateOne(userCommunicationService *model.UserCommunicationService) error

	DeleteOneByIDAndUserID(id uint8, userId uuid.UUID) error
}

type userCommunicationServiceRepository struct {
	db *gorm.DB
}

func NewUserCommunicationServiceRepository(db *gorm.DB) UserCommunicationServiceRepository {
	return &userCommunicationServiceRepository{db}
}

func (r *userCommunicationServiceRepository) FindManyByUserId(userId uuid.UUID) ([]model.UserCommunicationService, error) {
	var userCommunicationServices []model.UserCommunicationService
	if err := r.db.
		Model(&model.UserCommunicationService{}).
		Where("user_id = ?", userId).
		Find(&userCommunicationServices).Error; err != nil {
		return nil, err
	}
	return userCommunicationServices, nil
}

func (r *userCommunicationServiceRepository) FindManyByIdsAndUserId(ids []uint8, userId uuid.UUID) ([]model.UserCommunicationService, error) {
	if len(ids) == 0 {
		return []model.UserCommunicationService{}, nil
	}

	convertedIDs := make([]uint, len(ids))
	for i, id := range ids {
		convertedIDs[i] = uint(id)
	}

	var userCommunicationServices []model.UserCommunicationService
	if err := r.db.
		Model(&model.UserCommunicationService{}).
		Where("id IN ? AND user_id = ?", convertedIDs, userId).
		Find(&userCommunicationServices).Error; err != nil {
		return nil, err
	}
	return userCommunicationServices, nil
}

func (r *userCommunicationServiceRepository) FindOneByIDAndUserID(id uint8, userId uuid.UUID) (*model.UserCommunicationService, error) {
	var userCommunicationService model.UserCommunicationService
	if err := r.db.First(&userCommunicationService, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		return nil, err
	}
	return &userCommunicationService, nil
}

func (r *userCommunicationServiceRepository) CreateOne(userCommunicationService *model.UserCommunicationService) error {
	if err := r.db.Create(userCommunicationService).Error; err != nil {
		return err
	}
	return nil
}

func (r *userCommunicationServiceRepository) UpdateOne(userCommunicationService *model.UserCommunicationService) error {
	return r.db.
		Model(&model.UserCommunicationService{}).
		Where("id = ? AND user_id = ?", userCommunicationService.ID, userCommunicationService.UserId).
		Updates(map[string]interface{}{
			"name":    userCommunicationService.Name,
			"value":   userCommunicationService.Value,
			"service": userCommunicationService.Service,
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
