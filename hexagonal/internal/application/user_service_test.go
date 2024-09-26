package application_test

import (
	"errors"
	"hexagonal/internal/application"
	"hexagonal/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	mock.Mock
}

// Mock GetByID function
func (m *MockUserRepository) GetByID(id int) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

// Mock Save function
func (m *MockUserRepository) Save(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Test for GetUserByID
func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := application.NewUserService(mockRepo)

	user1 := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	user2 := &domain.User{ID: 2, Name: "Jane Doe", Email: "jane@example.com"}
	// Mock the behavior for the repository
	mockRepo.On("GetByID", 1).Return(user1, nil)
	mockRepo.On("GetByID", 2).Return(user2, nil)
	mockRepo.On("GetByID", 3).Return(nil, errors.New("user not found"))
	// Call the service method
	result1, err := userService.GetUserByID(1)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, result1)
	assert.Equal(t, "John Doe", result1.Name)
	assert.Equal(t, "john@example.com", result1.Email)

	result2, err := userService.GetUserByID(2)
	assert.Nil(t, err)
	assert.NotNil(t, result2)
	assert.Equal(t, "Jane Doe", result2.Name)
	assert.Equal(t, "jane@example.com", result2.Email)
	// Ensure the mock was called with the expected parameter

	result3, err := userService.GetUserByID(3)
	assert.NotNil(t, err)
	assert.Nil(t, result3)
	mockRepo.AssertExpectations(t)
}

func TestSaveUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := application.NewUserService(mockRepo)

	user := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	mockRepo.On("Save", user).Return(nil)
	err := userService.SaveUser(user)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
