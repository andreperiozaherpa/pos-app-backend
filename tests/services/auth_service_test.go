package services_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/core/services"
	"pos-app/backend/internal/models"
	"pos-app/backend/internal/utils"
	"pos-app/backend/tests/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_Login(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	authService := services.NewAuthService(mockUserRepo)

	password := "password123"
	hashedPassword, _ := utils.HashPassword(password)

	mockUser := &models.User{
		ID:           uuid.New(),
		Username:     sql.NullString{String: "testuser", Valid: true},
		PasswordHash: sql.NullString{String: hashedPassword, Valid: true},
		IsActive:     true,
	}

	t.Run("Success", func(t *testing.T) {
		// Setup mock
		mockUserRepo.On("GetByUsername", mock.Anything, mockUser.Username.String).Return(mockUser, nil).Once()

		// Execute
		user, token, err := authService.Login(context.Background(), mockUser.Username.String, password)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, token)
		assert.Equal(t, mockUser.ID, user.ID)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		// Setup mock
		mockUserRepo.On("GetByUsername", mock.Anything, "unknownuser").Return(nil, sql.ErrNoRows).Once()

		// Execute
		user, token, err := authService.Login(context.Background(), "unknownuser", "password123")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, services.ErrInvalidCredentials, err)
		assert.Nil(t, user)
		assert.Empty(t, token)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Incorrect Password", func(t *testing.T) {
		// Setup mock
		mockUserRepo.On("GetByUsername", mock.Anything, mockUser.Username.String).Return(mockUser, nil).Once()

		// Execute
		user, token, err := authService.Login(context.Background(), mockUser.Username.String, "wrongpassword")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, services.ErrInvalidCredentials, err)
		assert.Nil(t, user)
		assert.Empty(t, token)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Inactive User", func(t *testing.T) {
		inactiveUser := *mockUser // copy
		inactiveUser.IsActive = false

		// Setup mock
		mockUserRepo.On("GetByUsername", mock.Anything, inactiveUser.Username.String).Return(&inactiveUser, nil).Once()

		// Execute
		user, token, err := authService.Login(context.Background(), inactiveUser.Username.String, password)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, services.ErrUserInactive, err)
		assert.Nil(t, user)
		assert.Empty(t, token)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestAuthService_ValidateToken(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	authService := services.NewAuthService(mockUserRepo)

	mockUser := &models.User{
		ID:       uuid.New(),
		Username: sql.NullString{String: "validuser", Valid: true},
		IsActive: true,
	}

	// Generate a valid token for the test
	// In a real app, the secret and duration would come from config
	validToken, _ := utils.GenerateToken(mockUser.ID, "default_super_secret_key", time.Hour)

	t.Run("Success", func(t *testing.T) {
		// Setup mock
		mockUserRepo.On("GetByID", mock.Anything, mockUser.ID).Return(mockUser, nil).Once()

		// Execute
		user, err := authService.ValidateToken(context.Background(), validToken)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, mockUser.ID, user.ID)
		mockUserRepo.AssertExpectations(t)
	})
}
