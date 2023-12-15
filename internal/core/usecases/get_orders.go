package usecases

import "github.com/sousair/americastech-exchange/internal/core/entities"

type (
	GetOrdersParams struct {
		UserID string
	}

	GetOrdersUseCase interface {
		Get(params GetOrdersParams) ([]*entities.Order, error)
	}
)
