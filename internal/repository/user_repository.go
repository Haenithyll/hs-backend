package repository

import (
	"hs-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	ExistsById(id uuid.UUID) (bool, error)

	FindAll() ([]model.User, error)
	FindOneById(id uuid.UUID) (*model.User, error)
	FindOneByEmail(email string) (*model.User, error)
	FindOneByFingerprint(fingerprint string) (*model.User, error)
	FindManyByIds(uuids []uuid.UUID) ([]model.User, error)

	AssignKey(userId uuid.UUID, keyHash string, keyFingerprint string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) ExistsById(id uuid.UUID) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindOneById(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindOneByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindOneByFingerprint(fingerprint string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "api_key_fingerprint = ?", fingerprint).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindManyByIds(uuids []uuid.UUID) ([]model.User, error) {
	var users []model.User
	if err := r.db.Where("id IN ?", uuids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) AssignKey(userId uuid.UUID, keyHash string, keyFingerprint string) error {
	return r.db.Model(&model.User{}).
		Where("id = ?", userId).
		Updates(map[string]any{
			"api_key_hash":        keyHash,
			"api_key_fingerprint": keyFingerprint,
		}).Error
}
