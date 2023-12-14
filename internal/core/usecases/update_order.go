package usecases

import "github.com/sousair/americastech-exchange/internal/core/enums"

type (
	UpdateOrderParams struct {
		ExternalID string
		Price      string
		Status     enums.OrderStatus
	}

	UpdateOrderUseCase interface {
		Update(params UpdateOrderParams) error
	}
)
