package mocks

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// CustomerRepositoryMock adalah mock untuk postgres.CustomerRepository.
type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) Create(ctx context.Context, customer *models.Customer) error {
	args := m.Called(ctx, customer)
	return args.Error(0)
}

func (m *CustomerRepositoryMock) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Customer, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *CustomerRepositoryMock) GetByMembershipNumber(ctx context.Context, companyID uuid.UUID, membershipNumber string) (*models.Customer, error) {
	args := m.Called(ctx, companyID, membershipNumber)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *CustomerRepositoryMock) Update(ctx context.Context, customer *models.Customer) error {
	args := m.Called(ctx, customer)
	return args.Error(0)
}

func (m *CustomerRepositoryMock) Delete(ctx context.Context, userID uuid.UUID) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}
