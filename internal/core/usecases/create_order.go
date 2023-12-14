package usecases

import (
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/enums"
)

type (
	CreateOrderParams struct {
		UserID    string
		Pair      string
		Direction enums.OrderDirection
		Amount    string
		Type      enums.OrderType
		Price     string
	}

	CreateOrderUseCase interface {
		Create(params CreateOrderParams) (*entities.Order, error)
	}
)
