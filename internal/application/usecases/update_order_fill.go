package app_usecases

import (
	"fmt"

	custom_errors "github.com/sousair/americastech-exchange/internal/application/errors"
	"github.com/sousair/americastech-exchange/internal/application/providers/repositories"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type UpdateOrderFillUseCase struct {
	orderRepository repositories.OrderRepository
}

func NewUpdateOrderFillUseCase(orderRepository repositories.OrderRepository) usecases.UpdateOrderFillFieldsUseCase {
	return &UpdateOrderFillUseCase{
		orderRepository: orderRepository,
	}
}

func (uc UpdateOrderFillUseCase) Update(params usecases.UpdateOrderFillParams) error {
	order, err := uc.orderRepository.FindOneBy(map[string]interface{}{
		"external_id": params.ExternalID,
	})

	if err != nil {
		return err
	}

	if order == nil {
		return custom_errors.NewOrderNotFoundError(
			fmt.Errorf("order not found %s", params.ExternalID),
		)
	}

	if order.Status == enums.OrderStatusFilled {
		return custom_errors.NewOrderAlreadyFilledError(
			fmt.Errorf("order already filled %s", order.ID),
		)
	}

	if order.Status == enums.OrderStatusCanceled {
		return custom_errors.NewOrderAlreadyCanceledError(
			fmt.Errorf("order already canceled %s", order.ID),
		)
	}

	order.Status = enums.OrderStatus(params.Status)
	order.Price = params.Price

	_, err = uc.orderRepository.Update(order)

	if err != nil {
		return err
	}

	return nil
}
