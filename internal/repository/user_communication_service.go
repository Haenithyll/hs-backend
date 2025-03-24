package repository

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCommunicationServiceRepository interface {
	FindByUserId(userId uuid.UUID) ([]model.UserCommunicationService, error)
	FindByIDAndUserID(id uint8, userID uuid.UUID) (*model.UserCommunicationService, error)
	FindIDsByUserId(userId uuid.UUID) ([]uint8, error)

	Create(ucs *model.UserCommunicationService) error
	Update(ucs *model.UserCommunicationService) error

	DeleteByIDAndUserID(id uint8, userID uuid.UUID) error
}

type userCommunicationServiceRepository struct {
	db *gorm.DB
}

func NewUserCommunicationServiceRepository(db *gorm.DB) UserCommunicationServiceRepository {
	return &userCommunicationServiceRepository{db}
}

func (r *userCommunicationServiceRepository) FindByUserId(userId uuid.UUID) ([]model.UserCommunicationService, error) {
	var ucsList []model.UserCommunicationService
	if err := r.db.
		Model(&model.UserCommunicationService{}).
		Where("user_id = ?", userId).
		Find(&ucsList).Error; err != nil {
		return nil, err
	}
	return ucsList, nil
}

func (r *userCommunicationServiceRepository) FindByIDAndUserID(id uint8, userID uuid.UUID) (*model.UserCommunicationService, error) {
	var ucs model.UserCommunicationService
	if err := r.db.First(&ucs, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return nil, err
	}
	return &ucs, nil
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

func (r *userCommunicationServiceRepository) Create(ucs *model.UserCommunicationService) error {
	if err := r.db.Create(ucs).Error; err != nil {
		return err
	}
	return nil
}

func (r *userCommunicationServiceRepository) Update(ucs *model.UserCommunicationService) error {
	return r.db.
		Model(&model.UserCommunicationService{}).
		Where("id = ? AND user_id = ?", ucs.ID, ucs.UserId).
		Updates(map[string]interface{}{
			"name":    ucs.Name,
			"value":   ucs.Value,
			"service": ucs.Service,
		}).Error
}

func (r *userCommunicationServiceRepository) DeleteByIDAndUserID(id uint8, userID uuid.UUID) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.UserCommunicationService{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
