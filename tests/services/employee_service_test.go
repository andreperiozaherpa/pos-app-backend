package services_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"pos-app/backend/internal/core/services"
	"pos-app/backend/internal/models"
	"pos-app/backend/tests/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupEmployeeServiceTest() (services.EmployeeService, *mocks.UserRepositoryMock, *mocks.EmployeeRepositoryMock, *mocks.EmployeeRoleRepositoryMock, *mocks.RoleRepositoryMock) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	mockEmployeeRepo := new(mocks.EmployeeRepositoryMock)
	mockEmployeeRoleRepo := new(mocks.EmployeeRoleRepositoryMock)
	mockRoleRepo := new(mocks.RoleRepositoryMock)

	// Buat mock TransactionRunner. Ia hanya akan menjalankan fungsi logika bisnis
	// dengan mock repository yang sudah kita siapkan.
	mockRunner := func(ctx context.Context, fn services.TransactionFunc) error {
		return fn(mockUserRepo, mockEmployeeRepo, mockEmployeeRoleRepo, mockRoleRepo)
	}

	// Inisialisasi service dengan mock runner dan mock repositories.
	employeeService := services.NewEmployeeService(mockRunner, mockUserRepo, mockEmployeeRepo, mockEmployeeRoleRepo, mockRoleRepo)
	return employeeService, mockUserRepo, mockEmployeeRepo, mockEmployeeRoleRepo, mockRoleRepo
}

func TestEmployeeService_AddEmployee(t *testing.T) {
	employeeService, mockUserRepo, mockEmployeeRepo, mockEmployeeRoleRepo, mockRoleRepo := setupEmployeeServiceTest()

	companyID := uuid.New()
	storeID := uuid.New()
	roleID := int32(1)

	user := &models.User{
		Username: sql.NullString{String: "newemployee", Valid: true},
		Email:    sql.NullString{String: "newemployee@example.com", Valid: true},
	}
	employee := &models.Employee{
		CompanyID: companyID,
		StoreID:   uuid.NullUUID{UUID: storeID, Valid: true},
	}
	password := "password123"

	t.Run("Success", func(t *testing.T) {
		// Setup mocks
		mockRoleRepo.On("GetByID", mock.Anything, roleID).Return(&models.Role{ID: roleID}, nil).Once()
		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("GetByEmail", mock.Anything, user.Email.String).Return(nil, sql.ErrNoRows).Once()
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil).Once()
		mockEmployeeRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Employee")).Return(nil).Once()
		mockEmployeeRoleRepo.On("AssignRoleToEmployee", mock.Anything, mock.AnythingOfType("uuid.UUID"), roleID).Return(nil).Once()

		// Execute
		createdEmployee, err := employeeService.AddEmployee(context.Background(), employee, user, password, []int32{roleID})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, createdEmployee)
		assert.Equal(t, user.ID, createdEmployee.UserID) // UserID harus diisi
		mockRoleRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
		mockEmployeeRepo.AssertExpectations(t)
		mockEmployeeRoleRepo.AssertExpectations(t)
	})

	t.Run("Role Not Found", func(t *testing.T) {
		// Setup ulang service dengan mock baru untuk test case ini
		employeeService, _, _, _, mockRoleRepo := setupEmployeeServiceTest()
		// Reset user dan employee untuk memastikan tidak ada state dari test sebelumnya

		mockRoleRepo.On("GetByID", mock.Anything, roleID).Return(nil, sql.ErrNoRows).Once()

		_, err := employeeService.AddEmployee(context.Background(), employee, user, password, []int32{roleID})

		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("%w: dengan id %d", services.ErrRoleNotFound, roleID), err)
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("Username Exists", func(t *testing.T) {
		// Setup ulang service dengan mock baru
		employeeService, mockUserRepo, _, _, mockRoleRepo := setupEmployeeServiceTest()

		mockRoleRepo.On("GetByID", mock.Anything, roleID).Return(&models.Role{ID: roleID}, nil).Once()
		mockUserRepo.On("GetByUsername", mock.Anything, user.Username.String).Return(&models.User{}, nil).Once()

		_, err := employeeService.AddEmployee(context.Background(), employee, user, password, []int32{roleID})

		assert.Error(t, err)
		assert.Equal(t, services.ErrUsernameExists, err)
		mockRoleRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestEmployeeService_DeactivateEmployee(t *testing.T) {
	employeeService, mockUserRepo, _, _, _ := setupEmployeeServiceTest()
	employeeUserID := uuid.New()

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("Delete", employeeUserID).Return(nil).Once()

		err := employeeService.DeactivateEmployee(context.Background(), employeeUserID)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}
