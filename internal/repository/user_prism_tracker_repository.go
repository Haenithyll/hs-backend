package repository

import (
	"hs-backend/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserPrismTrackerRepository interface {
	FindAllWithPrisms() ([]model.UserPrismTracker, error)

	ActivatePrismByPrismIdAndUserId(prismId uint8, userId uuid.UUID) error
}

type userPrismTrackerRepository struct {
	db *gorm.DB
}

func NewUserPrismTrackerRepository(db *gorm.DB) UserPrismTrackerRepository {
	return &userPrismTrackerRepository{db}
}

func (r *userPrismTrackerRepository) FindAllWithPrisms() ([]model.UserPrismTracker, error) {
	var userPrismTrackers []model.UserPrismTracker
	if err := r.db.Preload("Prism").Find(&userPrismTrackers).Error; err != nil {
		return nil, err
	}
	return userPrismTrackers, nil
}

func (r *userPrismTrackerRepository) ActivatePrismByPrismIdAndUserId(prismId uint8, userId uuid.UUID) error {
	now := time.Now()

	tracker := model.UserPrismTracker{
		PrismID:       prismId,
		UserId:        userId,
		LastUpdatedAt: now,
	}

	return r.db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"prism_id", "last_updated_at"}),
		}).
		Create(&tracker).Error
}
