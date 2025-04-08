package user_test

// import (
// 	"errors"
// 	"net/http"
// 	"testing"

// 	"hs-backend/internal/handler/user"
// 	"hs-backend/internal/model"
// 	"hs-backend/internal/util"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// type MockUserRepository struct {
// 	FindOneByIdFunc func(userID uuid.UUID) (*model.User, error)
// }

// func (m *MockUserRepository) FindOneById(userID uuid.UUID) (*model.User, error) {
// 	return m.FindOneByIdFunc(userID)
// }

// func (m *MockUserRepository) FindOneByEmail(email string) (*model.User, error) {
// 	return nil, nil
// }

// func TestGetUserMeHandler_Nominal(t *testing.T) {
// 	mockUserId := uuid.New()
// 	c, w := util.NewTestGinContext("GET", "/api/users/me", "", mockUserId.String())

// 	mockRepo := &MockUserRepository{
// 		FindOneByIdFunc: func(userId uuid.UUID) (*model.User, error) {
// 			return &model.User{
// 				ID:        userId,
// 				FirstName: "John",
// 				LastName:  "Doe",
// 				Email:     "john.doe@example.com",
// 			}, nil
// 		},
// 	}

// 	handler := user.NewGetUserMeHandler(mockRepo)
// 	handler.Handle(c)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Contains(t, w.Body.String(), "John")
// }

// func TestGetUserMeHandler_MissingUser(t *testing.T) {
// 	mockUserId := uuid.New()
// 	c, w := util.NewTestGinContext("GET", "/api/users/me", "", mockUserId.String())

// 	mockRepo := &MockUserRepository{
// 		FindOneByIdFunc: func(userId uuid.UUID) (*model.User, error) {
// 			return nil, errors.New("user not found")
// 		},
// 	}

// 	handler := user.NewGetUserMeHandler(mockRepo)
// 	handler.Handle(c)

// 	assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	assert.Contains(t, w.Body.String(), "Failed to get user")
// }
