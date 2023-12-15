package usecases

import "github.com/sousair/americastech-exchange/internal/core/enums"

type (
	UpdateOrderFillParams struct {
		ExternalID string
		Price      string
		Status     enums.OrderStatus
	}

	UpdateOrderFillFieldsUseCase interface {
		Update(params UpdateOrderFillParams) error
	}
)
