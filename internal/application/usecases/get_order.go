package app_usecases

import (
	"github.com/sousair/americastech-exchange/internal/application/providers/repositories"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type GetOrderUseCase struct {
	orderRepository repositories.OrderRepository
}

func NewGetOrderUseCase(orderRepository repositories.OrderRepository) usecases.GetOrderUseCase {
	return &GetOrderUseCase{
		orderRepository: orderRepository,
	}
}

func (uc GetOrderUseCase) Get(params usecases.GetOrderParams) (*entities.Order, error) {
	order, err := uc.orderRepository.FindOneBy(map[string]interface{}{
		"id":     params.OrderId,
		"userId": params.UserId,
	})

	if err != nil {
		return nil, err
	}

	return order, nil
}
