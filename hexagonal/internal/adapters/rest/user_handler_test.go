package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"hexagonal/internal/adapters/rest"
	"hexagonal/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

// GetUserByID implements application.UserServicePort.
func (m *MockUserService) GetUserByID(id int) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

// SaveUser implements application.UserServicePort.
func (m *MockUserService) SaveUser(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	userHandler := rest.NewUserHandler(mockService)
	user1 := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	mockService.On("GetUserByID", 1).Return(user1, nil)
	mockService.On("GetUserByID", 2).Return(nil, errors.New("user not found"))
	r := gin.Default()
	r.GET("/user/:id", userHandler.GetByID)

	t.Run("Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/user/a", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		r.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusOK, w.Code)
	})
	t.Run("User 1", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/user/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp domain.User
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, "John Doe", resp.Name)
		assert.Equal(t, "john@example.com", resp.Email)
	})

	t.Run("User 1", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/user/2", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusOK, w.Code)
	})
	// Create Test Request

	mockService.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	userHandler := rest.NewUserHandler(mockService)
	user1 := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	user2 := &domain.User{ID: 2, Name: "Jane Doe", Email: "john@example.com"}
	mockService.On("SaveUser", user1).Return(nil)
	mockService.On("SaveUser", user2).Return(errors.New("email already exists"))
	r := gin.Default()
	r.POST("/user", userHandler.CreateUser)
	t.Run("Bad Request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/user", nil)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusCreated, w.Code)
	})

	t.Run("Create User 1", func(t *testing.T) {

		w := httptest.NewRecorder()
		userJSON, _ := json.Marshal(user1)
		req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userJSON))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("Create User 2", func(t *testing.T) {
		w := httptest.NewRecorder()
		userJSON, _ := json.Marshal(user2)
		req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userJSON))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusCreated, w.Code)
	})
	mockService.AssertExpectations(t)
}
