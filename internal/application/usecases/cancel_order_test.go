package app_usecases

import (
	"errors"
	"testing"

	custom_errors "github.com/sousair/americastech-exchange/internal/application/errors"
	exchange_mock "github.com/sousair/americastech-exchange/internal/application/providers/exchange/mocks"
	repositories_mock "github.com/sousair/americastech-exchange/internal/application/providers/repositories/mocks"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCancelOrderUseCase_CancelOrderRepositoryError(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	cancelOrderUC := NewCancelOrderUseCase(orderRepository, exchangeProvider)

	orderRepository.On("FindOneBy", mock.Anything).Return(nil, errors.New("error"))

	err := cancelOrderUC.Cancel(usecases.CancelOrderParams{
		OrderID: "1",
		UserID:  "1",
	})

	assert.Error(t, err)
}

func TestCancelOrderUseCase_CancelOrderNotFound(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	cancelOrderUC := NewCancelOrderUseCase(orderRepository, exchangeProvider)

	orderRepository.On("FindOneBy", mock.Anything).Return(nil, nil)

	err := cancelOrderUC.Cancel(usecases.CancelOrderParams{
		OrderID: "1",
		UserID:  "1",
	})

	assert.Error(t, err)
	assert.IsType(t, custom_errors.OrderNotFoundError, err)
}

func TestCancelOrderUseCase_CancelOrderAlreadyCanceled(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	cancelOrderUC := NewCancelOrderUseCase(orderRepository, exchangeProvider)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{
		Status: enums.OrderStatusCanceled,
	}, nil)

	err := cancelOrderUC.Cancel(usecases.CancelOrderParams{
		OrderID: "1",
		UserID:  "1",
	})

	assert.Error(t, err)
	assert.IsType(t, custom_errors.OrderAlreadyCanceledError, err)
}

func TestCancelOrderUseCase_CancelOrderAlreadyFilled(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	cancelOrderUC := NewCancelOrderUseCase(orderRepository, exchangeProvider)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{
		Status: enums.OrderStatusFilled,
	}, nil)

	err := cancelOrderUC.Cancel(usecases.CancelOrderParams{
		OrderID: "1",
		UserID:  "1",
	})

	assert.Error(t, err)
	assert.IsType(t, custom_errors.OrderAlreadyFilledError, err)
}

func TestCancelOrderUseCase_CancelExchangeError(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	cancelOrderUC := NewCancelOrderUseCase(orderRepository, exchangeProvider)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{
		Status: enums.OrderStatusPartiallyFilled,
	}, nil)

	exchangeProvider.On("CancelOrder", mock.Anything).Return(errors.New("error"))

	err := cancelOrderUC.Cancel(usecases.CancelOrderParams{
		OrderID: "1",
		UserID:  "1",
	})

	assert.Error(t, err)
}

func TestCancelOrderUseCase_CancelOrderUpdateError(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	cancelOrderUC := NewCancelOrderUseCase(orderRepository, exchangeProvider)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{
		Status: enums.OrderStatusPartiallyFilled,
	}, nil)

	exchangeProvider.On("CancelOrder", mock.Anything).Return(nil)

	orderRepository.On("Update", mock.Anything).Return(nil, errors.New("error"))

	err := cancelOrderUC.Cancel(usecases.CancelOrderParams{
		OrderID: "1",
		UserID:  "1",
	})

	assert.Error(t, err)
}

func TestCancelOrderUseCase_CancelSuccess(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)
	exchangeProvider := new(exchange_mock.MockExchangeProvider)

	cancelOrderUC := NewCancelOrderUseCase(orderRepository, exchangeProvider)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{
		Status: enums.OrderStatusPartiallyFilled,
	}, nil)

	exchangeProvider.On("CancelOrder", mock.Anything).Return(nil)

	orderRepository.On("Update", mock.Anything).Return(&entities.Order{}, nil)

	err := cancelOrderUC.Cancel(usecases.CancelOrderParams{
		OrderID: "1",
		UserID:  "1",
	})

	assert.NoError(t, err)
}
