package app_usecases

import (
	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	"github.com/sousair/americastech-exchange/internal/application/providers/repositories"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type CreateOrderUseCase struct {
	orderRepository  repositories.OrderRepository
	exchangeProvider exchange.ExchangeProvider
}

func NewCreateOrderUseCase(orderRepository repositories.OrderRepository, exchangeProvider exchange.ExchangeProvider) usecases.CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepository:  orderRepository,
		exchangeProvider: exchangeProvider,
	}
}

func (uc CreateOrderUseCase) Create(params usecases.CreateOrderParams) (*entities.Order, error) {
	createdOrder, err := uc.exchangeProvider.Create(exchange.CreateOrderParams{
		Pair:      params.Pair,
		Direction: params.Direction,
		Amount:    params.Amount,
		Type:      params.Type,
		Price:     params.Price,
	})

	if err != nil {
		return nil, err
	}

	order := &entities.Order{
		ExternalID: createdOrder.ExternalID,
		UserID:     params.UserID,
		Pair:       createdOrder.Pair,
		Direction:  createdOrder.Direction,
		Amount:     createdOrder.Amount,
		Type:       createdOrder.Type,
		Price:      createdOrder.Price,
		Status:     createdOrder.Status,
	}

	order, err = uc.orderRepository.Create(order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
