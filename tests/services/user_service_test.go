package services_test

import (
	"context"
	"database/sql"
	"testing"

	"pos-app/backend/internal/core/services"
	"pos-app/backend/internal/models"
	"pos-app/backend/internal/utils"
	"pos-app/backend/tests/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_RegisterUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	userService := services.NewUserService(mockUserRepo)

	t.Run("Success", func(t *testing.T) {
		user := &models.User{
			UserType: models.UserTypeEmployee,
			Username: sql.NullString{String: "newuser", Valid: true},
			Email:    sql.NullString{String: "newuser@example.com", Valid: true},
		}
		password := "securepassword"

		// Mock GetByUsername dan GetByEmail untuk mengembalikan sql.ErrNoRows (tidak ditemukan)
		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("GetByEmail", mock.Anything, user.Email.String).Return(nil, sql.ErrNoRows).Once()
		// Mock Create untuk sukses
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil).Once()

		err := userService.RegisterUser(context.Background(), user, password)

		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, user.ID) // ID harus di-generate
		assert.True(t, user.PasswordHash.Valid)
		assert.True(t, utils.CheckPasswordHash(password, user.PasswordHash.String)) // Password harus di-hash
		assert.True(t, user.IsActive)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Username Already Exists", func(t *testing.T) {
		existingUser := &models.User{Username: sql.NullString{String: "existinguser", Valid: true}}
		user := &models.User{Username: sql.NullString{String: "existinguser", Valid: true}}
		password := "password"

		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(existingUser, nil).Once()

		err := userService.RegisterUser(context.Background(), user, password)

		assert.Error(t, err)
		assert.Equal(t, services.ErrUsernameExists, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Email Already Exists", func(t *testing.T) {
		existingUser := &models.User{Email: sql.NullString{String: "existing@example.com", Valid: true}}
		user := &models.User{
			Username: sql.NullString{String: "anotheruser", Valid: true},
			Email:    sql.NullString{String: "existing@example.com", Valid: true},
		}
		password := "password"

		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("GetByEmail", mock.Anything, user.Email.String).Return(existingUser, nil).Once()

		err := userService.RegisterUser(context.Background(), user, password)

		assert.Error(t, err)
		assert.Equal(t, services.ErrEmailExists, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUserProfile(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	userService := services.NewUserService(mockUserRepo)

	mockUser := &models.User{
		ID:       uuid.New(),
		Username: sql.NullString{String: "testuser", Valid: true},
	}

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mockUser.ID).Return(mockUser, nil).Once()

		foundUser, err := userService.GetUserProfile(context.Background(), mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, mockUser.ID, foundUser.ID)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		nonExistentID := uuid.New()
		mockUserRepo.On("GetByID", mock.Anything, nonExistentID).Return(nil, sql.ErrNoRows).Once()

		foundUser, err := userService.GetUserProfile(context.Background(), nonExistentID)

		assert.Error(t, err)
		assert.Equal(t, services.ErrUserNotFound, err)
		assert.Nil(t, foundUser)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_UpdateUserProfile(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	userService := services.NewUserService(mockUserRepo)

	mockUser := &models.User{
		ID:       uuid.New(),
		Username: sql.NullString{String: "testuser", Valid: true},
		FullName: sql.NullString{String: "Old Name", Valid: true},
	}

	t.Run("Success", func(t *testing.T) {
		updatedUser := *mockUser // Copy
		updatedUser.FullName = sql.NullString{String: "New Name", Valid: true}

		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil).Once()

		err := userService.UpdateUserProfile(context.Background(), &updatedUser)

		assert.NoError(t, err)
		// Verify that UpdatedAt was set
		assert.True(t, updatedUser.UpdatedAt.After(mockUser.UpdatedAt))
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("User Not Found During Update", func(t *testing.T) {
		nonExistentUser := &models.User{
			ID:       uuid.New(),
			Username: sql.NullString{String: "nonexistent", Valid: true},
		}

		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.User")).Return(sql.ErrNoRows).Once()

		err := userService.UpdateUserProfile(context.Background(), nonExistentUser)

		assert.Error(t, err)
		assert.Equal(t, services.ErrUserNotFound, err)
		mockUserRepo.AssertExpectations(t)
	})
}
