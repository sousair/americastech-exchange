package repositories_mock

import (
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Create(order *entities.Order) (*entities.Order, error) {
	args := m.Called(order)

	orderR := args.Get(0)
	if orderR == nil {
		return nil, args.Error(1)
	}

	return orderR.(*entities.Order), args.Error(1)
}

func (m *MockOrderRepository) FindOneBy(where map[string]interface{}) (*entities.Order, error) {
	args := m.Called(where)

	orderR := args.Get(0)
	if orderR == nil {
		return nil, args.Error(1)
	}

	return orderR.(*entities.Order), args.Error(1)
}

func (m *MockOrderRepository) FindAllBy(where map[string]interface{}) ([]*entities.Order, error) {
	args := m.Called(where)

	orderR := args.Get(0)
	if orderR == nil {
		return nil, args.Error(1)
	}

	return orderR.([]*entities.Order), args.Error(1)
}

func (m *MockOrderRepository) Update(params *entities.Order) (*entities.Order, error) {
	args := m.Called(params)

	orderR := args.Get(0)
	if orderR == nil {
		return nil, args.Error(1)
	}

	return orderR.(*entities.Order), args.Error(1)
}
