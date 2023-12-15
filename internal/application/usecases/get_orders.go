package app_usecases

import (
	"fmt"

	custom_errors "github.com/sousair/americastech-exchange/internal/application/errors"
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

	if orders == nil {
		return nil, custom_errors.NewOrderNotFoundError(
			fmt.Errorf("orders not found %s", params.UserID),
		)
	}

	return orders, nil
}
