package usecases

import "github.com/sousair/americastech-exchange/internal/core/entities"

type (
	GetOrderParams struct {
		UserID string
	}

	GetOrdersUseCase interface {
		Get(params GetOrderParams) ([]*entities.Order, error)
	}
)
