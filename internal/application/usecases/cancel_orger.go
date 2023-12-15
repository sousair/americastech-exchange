package app_usecases

import (
	"fmt"

	custom_errors "github.com/sousair/americastech-exchange/internal/application/errors"
	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	"github.com/sousair/americastech-exchange/internal/application/providers/repositories"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type CancelOrderUseCase struct {
	orderRepository  repositories.OrderRepository
	exchangeProvider exchange.ExchangeProvider
}

func NewCancelOrderUseCase(orderRepository repositories.OrderRepository, exchangeProvider exchange.ExchangeProvider) usecases.CancelOrderUseCase {
	return &CancelOrderUseCase{
		orderRepository:  orderRepository,
		exchangeProvider: exchangeProvider,
	}
}

func (c CancelOrderUseCase) Cancel(params usecases.CancelOrderParams) error {
	order, err := c.orderRepository.FindOneBy(map[string]interface{}{
		"id":      params.OrderID,
		"user_id": params.UserID,
	})

	if err != nil {
		return err
	}

	if order == nil {
		return custom_errors.NewOrderNotFoundError(
			fmt.Errorf("order not found %s", params.OrderID),
		)
	}

	if order.Status == enums.OrderStatusCanceled {
		return custom_errors.NewOrderAlreadyCanceledError(
			fmt.Errorf("order already canceled %s", params.OrderID),
		)
	}

	if order.Status == enums.OrderStatusFilled {
		return custom_errors.NewOrderAlreadyFilledError(
			fmt.Errorf("order already filled %s", params.OrderID),
		)
	}

	err = c.exchangeProvider.CancelOrder(order)

	if err != nil {
		return err
	}

	order.Status = enums.OrderStatusCanceled

	_, err = c.orderRepository.Update(order)

	if err != nil {
		return err
	}

	return nil
}
