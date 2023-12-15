package app_usecases

import (
	"github.com/sousair/americastech-exchange/internal/application/providers/repositories"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type GetOrdersUseCase struct {
	orderRepository repositories.OrderRepository
}

func NewGetOrdersUseCase(orderRepository repositories.OrderRepository) usecases.GetOrdersUseCase {
	return &GetOrdersUseCase{
		orderRepository: orderRepository,
	}
}

func (uc GetOrdersUseCase) Get(params usecases.GetOrdersParams) ([]*entities.Order, error) {
	orders, err := uc.orderRepository.FindAllBy(map[string]interface{}{
		"user_id": params.UserID,
	})

	if err != nil {
		return nil, err
	}

	return orders, nil
}
