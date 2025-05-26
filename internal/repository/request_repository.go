package repository

import (
	"hs-backend/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RequestRepository interface {
	FindManyEnrichedByReceiverId(receiverId uuid.UUID) ([]model.Request, error)
	FindManyEnrichedByIssuerId(issuerId uuid.UUID) ([]model.Request, error)
	FindOneByIdAndReceiverId(id uint8, receiverId uuid.UUID) (*model.Request, error)
	FindOneEnrichedById(id uint8) (*model.Request, error)

	CreateOne(request *model.Request) error
	MarkOneAsRead(requestId uint8) error
}

type requestRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) RequestRepository {
	return &requestRepository{db}
}

func (r *requestRepository) FindManyEnrichedByReceiverId(receiverId uuid.UUID) ([]model.Request, error) {
	var requests []model.Request
	if err := r.db.
		Preload("Issuer").
		Preload("Receiver").
		Preload("Level").
		Where("receiver_id = ?", receiverId).
		Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *requestRepository) FindManyEnrichedByIssuerId(issuerId uuid.UUID) ([]model.Request, error) {
	var requests []model.Request
	if err := r.db.
		Preload("Issuer").
		Preload("Receiver").
		Preload("Level").
		Where("issuer_id = ?", issuerId).
		Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *requestRepository) FindOneByIdAndReceiverId(id uint8, receiverId uuid.UUID) (*model.Request, error) {
	var request model.Request
	if err := r.db.Model(&model.Request{}).
		Where("id = ?", id).
		Where("receiver_id = ?", receiverId).
		First(&request).Error; err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *requestRepository) FindOneEnrichedById(id uint8) (*model.Request, error) {
	var request model.Request
	if err := r.db.
		Preload("Issuer").
		Preload("Receiver").
		Preload("Level").
		Where("id = ?", id).
		First(&request).Error; err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *requestRepository) CreateOne(request *model.Request) error {
	return r.db.Create(request).Error
}

func (r *requestRepository) MarkOneAsRead(requestId uint8) error {
	now := time.Now()

	return r.db.Model(&model.Request{}).
		Where("id = ?", requestId).
		Updates(map[string]any{
			"is_read": true,
			"read_at": now,
		}).Error
}
