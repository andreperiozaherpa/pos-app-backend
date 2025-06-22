package services_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"pos-app/backend/internal/core/services"
	"pos-app/backend/internal/models"
	"pos-app/backend/internal/utils"
	"pos-app/backend/tests/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupCustomerServiceTest() (services.CustomerService, *mocks.UserRepositoryMock, *mocks.CustomerRepositoryMock) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	mockCustomerRepo := new(mocks.CustomerRepositoryMock)

	// Buat mock CustomerTransactionRunner. Ia hanya akan menjalankan fungsi logika bisnis
	// dengan mock repository yang sudah kita siapkan.
	mockRunner := func(ctx context.Context, fn services.CustomerTransactionFunc) error {
		return fn(mockUserRepo, mockCustomerRepo)
	}

	// Inisialisasi service dengan mock runner dan mock repositories.
	customerService := services.NewCustomerService(mockRunner, mockUserRepo, mockCustomerRepo)
	return customerService, mockUserRepo, mockCustomerRepo
}

func TestCustomerService_AddCustomer(t *testing.T) {
	customerService, mockUserRepo, mockCustomerRepo := setupCustomerServiceTest()

	companyID := uuid.New()
	user := &models.User{
		Username:    sql.NullString{String: "newcustomer", Valid: true},
		Email:       sql.NullString{String: "newcustomer@example.com", Valid: true},
		PhoneNumber: sql.NullString{String: "081234567890", Valid: true},
	}
	customer := &models.Customer{
		CompanyID:        companyID,
		MembershipNumber: sql.NullString{String: "MEM-001", Valid: true},
	}
	password := "customerpass"

	t.Run("Success", func(t *testing.T) {
		// Setup mocks
		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("GetByEmail", mock.Anything, user.Email.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("GetByPhoneNumber", mock.Anything, user.PhoneNumber.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil).Once()
		mockCustomerRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Customer")).Return(nil).Once()

		// Execute
		createdCustomer, err := customerService.AddCustomer(context.Background(), customer, user, password)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, createdCustomer)
		assert.NotEqual(t, uuid.Nil, createdCustomer.UserID)
		assert.True(t, utils.CheckPasswordHash(password, user.PasswordHash.String))
		mockUserRepo.AssertExpectations(t)
		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("Username Exists", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(&models.User{}, nil).Once()

		_, err := customerService.AddCustomer(context.Background(), customer, user, password)

		assert.Error(t, err)
		assert.Equal(t, services.ErrUsernameExists, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Email Exists", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("GetByEmail", mock.Anything, user.Email.String).Return(&models.User{}, nil).Once()

		_, err := customerService.AddCustomer(context.Background(), customer, user, password)

		assert.Error(t, err)
		assert.Equal(t, services.ErrEmailExists, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Phone Number Exists", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("GetByEmail", mock.Anything, user.Email.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("GetByPhoneNumber", mock.Anything, user.PhoneNumber.String).Return(&models.User{}, nil).Once()

		_, err := customerService.AddCustomer(context.Background(), customer, user, password)

		assert.Error(t, err)
		assert.Equal(t, services.ErrPhoneNumberExists, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestCustomerService_UpdateCustomer(t *testing.T) {
	customerService, mockUserRepo, mockCustomerRepo := setupCustomerServiceTest()

	customerUserID := uuid.New()
	mockUser := &models.User{ID: customerUserID, FullName: sql.NullString{String: "Old Name", Valid: true}}
	mockCustomer := &models.Customer{UserID: customerUserID, Points: 100}

	t.Run("Success", func(t *testing.T) {
		updatedUser := *mockUser
		updatedUser.FullName = sql.NullString{String: "New Name", Valid: true}
		updatedCustomer := *mockCustomer
		updatedCustomer.Tier = sql.NullString{String: "Gold", Valid: true}

		mockUserRepo.On("Update", mock.Anything, &updatedUser).Return(nil).Once()
		mockCustomerRepo.On("Update", mock.Anything, &updatedCustomer).Return(nil).Once()

		err := customerService.UpdateCustomer(context.Background(), &updatedCustomer, &updatedUser)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("User Not Found During Update", func(t *testing.T) {
		updatedUser := *mockUser
		updatedCustomer := *mockCustomer

		mockUserRepo.On("Update", mock.Anything, &updatedUser).Return(sql.ErrNoRows).Once()

		err := customerService.UpdateCustomer(context.Background(), &updatedCustomer, &updatedUser)

		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err) // UserRepo returns sql.ErrNoRows directly
		mockUserRepo.AssertExpectations(t)
		mockCustomerRepo.AssertNotCalled(t, "Update") // CustomerRepo should not be called
	})
}

func TestCustomerService_GetCustomerByPhoneNumber(t *testing.T) {
	customerService, mockUserRepo, mockCustomerRepo := setupCustomerServiceTest()

	phoneNumber := "081234567890"
	customerUserID := uuid.New()
	mockUser := &models.User{ID: customerUserID, PhoneNumber: sql.NullString{String: phoneNumber, Valid: true}}
	mockCustomer := &models.Customer{UserID: customerUserID, CompanyID: uuid.New()}

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("GetByPhoneNumber", mock.Anything, phoneNumber).Return(mockUser, nil).Once()
		mockCustomerRepo.On("GetByUserID", mock.Anything, customerUserID).Return(mockCustomer, nil).Once()

		foundCustomer, err := customerService.GetCustomerByPhoneNumber(context.Background(), phoneNumber)

		assert.NoError(t, err)
		assert.NotNil(t, foundCustomer)
		assert.Equal(t, customerUserID, foundCustomer.UserID)
		assert.Equal(t, mockUser.PhoneNumber.String, foundCustomer.User.PhoneNumber.String) // Check if user is attached
		mockUserRepo.AssertExpectations(t)
		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("Phone Number Not Found", func(t *testing.T) {
		mockUserRepo.On("GetByPhoneNumber", mock.Anything, phoneNumber).Return(nil, sql.ErrNoRows).Once()

		_, err := customerService.GetCustomerByPhoneNumber(context.Background(), phoneNumber)

		assert.Error(t, err)
		assert.Equal(t, services.ErrCustomerNotFound, err)
		mockUserRepo.AssertExpectations(t)
		mockCustomerRepo.AssertNotCalled(t, "GetByUserID")
	})
}

func TestCustomerService_UpdateCustomerPoints(t *testing.T) {
	customerUserID := uuid.New()
	initialPoints := 100
	newPoints := 250

	t.Run("Success", func(t *testing.T) {
		// Setup di dalam sub-test untuk isolasi yang lebih baik
		customerService, _, mockCustomerRepo := setupCustomerServiceTest() // Deklarasi lokal
		mockCustomer := &models.Customer{UserID: customerUserID, Points: int32(initialPoints)}

		mockCustomerRepo.On("GetByUserID", mock.Anything, customerUserID).Return(mockCustomer, nil).Once()
		mockCustomerRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.Customer")).Return(nil).Once()

		err := customerService.UpdateCustomerPoints(context.Background(), customerUserID, newPoints)

		assert.NoError(t, err)
		// Bandingkan dengan tipe yang benar (int32)
		assert.Equal(t, int32(newPoints), mockCustomer.Points) // Check if points were updated in the object passed to repo
		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("Customer Not Found", func(t *testing.T) {
		// Setup ulang service dengan mock baru untuk test case ini
		customerService, _, mockCustomerRepo := setupCustomerServiceTest() // Deklarasi lokal

		mockCustomerRepo.On("GetByUserID", mock.Anything, customerUserID).Return(nil, sql.ErrNoRows).Once()

		err := customerService.UpdateCustomerPoints(context.Background(), customerUserID, newPoints)

		assert.Error(t, err)
		assert.Equal(t, services.ErrCustomerNotFound, err)
		mockCustomerRepo.AssertExpectations(t)
		mockCustomerRepo.AssertNotCalled(t, "Update")
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Setup ulang service dengan mock baru untuk test case ini
		customerService, _, mockCustomerRepo := setupCustomerServiceTest() // Deklarasi lokal
		mockCustomer := &models.Customer{UserID: customerUserID, Points: int32(initialPoints)}

		mockCustomerRepo.On("GetByUserID", mock.Anything, customerUserID).Return(mockCustomer, nil).Once()
		mockCustomerRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.Customer")).Return(fmt.Errorf("db error")).Once()

		err := customerService.UpdateCustomerPoints(context.Background(), customerUserID, newPoints)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db error")
		mockCustomerRepo.AssertExpectations(t)
	})
}
