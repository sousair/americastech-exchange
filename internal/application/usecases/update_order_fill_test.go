package app_usecases

import (
	"errors"
	"testing"

	custom_errors "github.com/sousair/americastech-exchange/internal/application/errors"
	repositories_mock "github.com/sousair/americastech-exchange/internal/application/providers/repositories/mocks"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateOrderFillUseCase_UpdateOrderRepositoryFindError(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)

	updateOrderFillUC := NewUpdateOrderFillUseCase(orderRepository)

	orderRepository.On("FindOneBy", mock.Anything).Return(nil, errors.New("error"))

	err := updateOrderFillUC.Update(usecases.UpdateOrderFillParams{})

	assert.Error(t, err)
}

func TestUpdateOrderFillUseCase_UpdateOrderRepositoryFindOrderNotFound(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)

	updateOrderFillUC := NewUpdateOrderFillUseCase(orderRepository)

	orderRepository.On("FindOneBy", mock.Anything).Return(nil, nil)

	err := updateOrderFillUC.Update(usecases.UpdateOrderFillParams{})

	assert.Error(t, err)
	assert.IsType(t, err, custom_errors.OrderNotFoundError)
}

func TestUpdateOrderFillUseCase_UpdateOrderRepositoryFindOrderAlreadyFilled(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)

	updateOrderFillUC := NewUpdateOrderFillUseCase(orderRepository)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{
		Status: enums.OrderStatusFilled,
	}, nil)

	err := updateOrderFillUC.Update(usecases.UpdateOrderFillParams{})

	assert.Error(t, err)
	assert.IsType(t, err, custom_errors.OrderAlreadyFilledError)
}

func TestUpdateOrderFillUseCase_UpdateOrderRepositoryFindOrderAlreadyCanceled(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)

	updateOrderFillUC := NewUpdateOrderFillUseCase(orderRepository)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{
		Status: enums.OrderStatusCanceled,
	}, nil)

	err := updateOrderFillUC.Update(usecases.UpdateOrderFillParams{})

	assert.Error(t, err)
	assert.IsType(t, err, custom_errors.OrderAlreadyCanceledError)
}

func TestUpdateOrderFillUseCase_UpdateOrderRepositoryUpdateError(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)

	updateOrderFillUC := NewUpdateOrderFillUseCase(orderRepository)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{}, nil)

	orderRepository.On("Update", mock.Anything).Return(nil, errors.New("error"))

	err := updateOrderFillUC.Update(usecases.UpdateOrderFillParams{})

	assert.Error(t, err)
}

func TestUpdateOrderFillUseCase_UpdateSuccess(t *testing.T) {
	orderRepository := new(repositories_mock.MockOrderRepository)

	updateOrderFillUC := NewUpdateOrderFillUseCase(orderRepository)

	orderRepository.On("FindOneBy", mock.Anything).Return(&entities.Order{}, nil)

	orderRepository.On("Update", mock.Anything).Return(&entities.Order{}, nil)

	err := updateOrderFillUC.Update(usecases.UpdateOrderFillParams{})

	assert.NoError(t, err)
}
