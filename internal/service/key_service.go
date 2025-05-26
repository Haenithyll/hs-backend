package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"hs-backend/internal/domain"
	"hs-backend/internal/repository"
	"hs-backend/internal/response"
	"hs-backend/internal/response/mapper"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type KeyService struct {
	UserRepository repository.UserRepository
}

func NewKeyService(userRepository repository.UserRepository) *KeyService {
	return &KeyService{UserRepository: userRepository}
}

func (s *KeyService) GenerateKey(userId uuid.UUID) (*response.KeyResponse, *domain.DomainError) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to generate key: "+err.Error())
	}

	key := hex.EncodeToString(bytes)

	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to generate key: "+err.Error())
	}

	keyHash := string(hash)

	sum := sha256.Sum256([]byte(key))
	keyFingerprint := hex.EncodeToString(sum[:])

	err = s.UserRepository.AssignKey(userId, keyHash, keyFingerprint)
	if err != nil {
		return nil, domain.NewDomainError(domain.ErrInternalServerError, "failed to assign key: "+err.Error())
	}

	return mapper.ToKeyResponse(key), nil
}
