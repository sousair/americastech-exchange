package app_usecases

import (
	"errors"
	"testing"

	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	exchange_mock "github.com/sousair/americastech-exchange/internal/application/providers/exchange/mocks"
	repositories_mock "github.com/sousair/americastech-exchange/internal/application/providers/repositories/mocks"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrderUseCase_CreateExchangeProviderError(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	createOrderUC := NewCreateOrderUseCase(orderRepository, exchangeProvider)

	exchangeProvider.On("Create", mock.Anything).Return(nil, errors.New("error"))

	order, err := createOrderUC.Create(usecases.CreateOrderParams{})

	assert.Error(t, err)
	assert.Nil(t, order)
}
func TestCreateOrderUseCase_CreateOrderRepositoryError(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	createOrderUC := NewCreateOrderUseCase(orderRepository, exchangeProvider)

	exchangeProvider.On("Create", mock.Anything).Return(&exchange.CreatedOrder{
		ExternalID: "1",
	}, nil)
	orderRepository.On("Create", mock.Anything).Return(nil, errors.New("error"))

	order, err := createOrderUC.Create(usecases.CreateOrderParams{})

	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestCreateOrderUseCase_CreateSuccess(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	createOrderUC := NewCreateOrderUseCase(orderRepository, exchangeProvider)

	exchangeProvider.On("Create", mock.Anything).Return(&exchange.CreatedOrder{
		ExternalID: "1",
	}, nil)
	orderRepository.On("Create", mock.Anything).Return(&entities.Order{
		ExternalID: "1",
	}, nil)

	order, err := createOrderUC.Create(usecases.CreateOrderParams{})

	assert.NoError(t, err)
	assert.NotNil(t, order)
}
