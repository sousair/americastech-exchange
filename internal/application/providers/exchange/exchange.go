package exchange

import (
	"github.com/sousair/americastech-exchange/internal/core/enums"
)

type (
	CreateOrderParams struct {
		Pair      string
		Direction enums.OrderDirection
		Type      enums.OrderType
		Amount    string
		Price     string
	}

	CreatedOrder struct {
		ExternalID string
		Pair       string
		Direction  enums.OrderDirection
		Type       enums.OrderType
		Amount     string
		Price      string
		Status     enums.OrderStatus
	}

	ExchangeProvider interface {
		Create(params CreateOrderParams) (*CreatedOrder, error)
	}
)
