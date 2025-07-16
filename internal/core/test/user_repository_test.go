package test

import (
	"context"
	"testing"

	"pos-app/backend/internal/core/mock"
	"pos-app/backend/internal/data/postgres"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()
	dbMock := new(mock.DBMock)
	repo := postgres.NewUserRepositoryPG(dbMock)

	user := &models.User{
		ID:       uuid.New(),
		Username: "testuser",
		Email:    "test@example.com",
		// Tambahkan field lain sesuai struct User
	}

	// Setup expectation: ExecContext dipanggil dan mengembalikan nil error
	dbMock.On("ExecContext", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

	err := repo.Create(ctx, user)
	assert.NoError(t, err)

	dbMock.AssertExpectations(t)
}

func TestUserRepository_GetByID(t *testing.T) {
	ctx := context.Background()
	dbMock := new(mock.DBMock)
	repo := postgres.NewUserRepositoryPG(dbMock)

	userID := uuid.New()
	expectedUser := &models.User{
		ID:       userID,
		Username: "testuser",
		Email:    "test@example.com",
		// Lengkapi sesuai kebutuhan
	}

	// Setup expectation: QueryRowContext dipanggil dan mengembalikan RowMock dengan ScanFunc mengisi expectedUser
	dbMock.On("QueryRowContext", mock.Anything, mock.Anything, mock.Anything).Return(&mock.RowMock{
		ScanFunc: func(dest ...interface{}) error {
			*(dest[0].(*uuid.UUID)) = expectedUser.ID
			*(dest[1].(*string)) = expectedUser.Username
			*(dest[2].(*string)) = expectedUser.Email
			// Isi field lain sesuai urutan Scan di repo
			return nil
		},
	}).Once()

	user, err := repo.GetByID(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Equal(t, expectedUser.Email, user.Email)

	dbMock.AssertExpectations(t)
}

func TestUserRepository_GetByUsername(t *testing.T) {
	ctx := context.Background()
	dbMock := new(mock.DBMock)
	repo := postgres.NewUserRepositoryPG(dbMock)

	username := "testuser"
	expectedUser := &models.User{
		ID:       uuid.New(),
		Username: username,
		Email:    "test@example.com",
	}

	dbMock.On("QueryRowContext", mock.Anything, mock.Anything, mock.Anything).Return(&mock.RowMock{
		ScanFunc: func(dest ...interface{}) error {
			*(dest[0].(*uuid.UUID)) = expectedUser.ID
			*(dest[1].(*string)) = expectedUser.Username
			*(dest[2].(*string)) = expectedUser.Email
			return nil
		},
	}).Once()

	user, err := repo.GetByUsername(ctx, username)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Equal(t, expectedUser.Email, user.Email)

	dbMock.AssertExpectations(t)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	ctx := context.Background()
	dbMock := new(mock.DBMock)
	repo := postgres.NewUserRepositoryPG(dbMock)

	email := "test@example.com"
	expectedUser := &models.User{
		ID:       uuid.New(),
		Username: "testuser",
		Email:    email,
	}

	dbMock.On("QueryRowContext", mock.Anything, mock.Anything, mock.Anything).Return(&mock.RowMock{
		ScanFunc: func(dest ...interface{}) error {
			*(dest[0].(*uuid.UUID)) = expectedUser.ID
			*(dest[1].(*string)) = expectedUser.Username
			*(dest[2].(*string)) = expectedUser.Email
			return nil
		},
	}).Once()

	user, err := repo.GetByEmail(ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Equal(t, expectedUser.Email, user.Email)

	dbMock.AssertExpectations(t)
}
