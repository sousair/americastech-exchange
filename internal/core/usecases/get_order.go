package usecases

import "github.com/sousair/americastech-exchange/internal/core/entities"

type (
	GetOrderParams struct {
		OrderId string
		UserId  string
	}

	GetOrderUseCase interface {
		Get(params GetOrderParams) (*entities.Order, error)
	}
)
