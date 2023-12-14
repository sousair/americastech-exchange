package app_usecases

import (
	"github.com/sousair/americastech-exchange/internal/application/providers/repositories"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type UpdateOrderUseCase struct {
	orderRepository repositories.OrderRepository
}

func NewUpdateOrderUseCase(orderRepository repositories.OrderRepository) usecases.UpdateOrderUseCase {
	return &UpdateOrderUseCase{
		orderRepository: orderRepository,
	}
}

func (uc UpdateOrderUseCase) Update(params usecases.UpdateOrderParams) error {
	order, err := uc.orderRepository.FindOneBy(map[string]interface{}{
		"external_id": params.ExternalID,
	})

	if err != nil {
		return err
	}

	order.Status = enums.OrderStatus(params.Status)
	order.Price = params.Price

	if err != nil {
		return err
	}

	return nil
}
