package exchange_mock

import (
	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/stretchr/testify/mock"
)

type MockExchangeProvider struct {
	mock.Mock
}

func (m *MockExchangeProvider) CancelOrder(order *entities.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockExchangeProvider) Create(params exchange.CreateOrderParams) (*exchange.CreatedOrder, error) {
	args := m.Called(params)

	createOrderR := args.Get(0)

	if createOrderR == nil {
		return nil, args.Error(1)
	}

	return createOrderR.(*exchange.CreatedOrder), args.Error(1)
}
