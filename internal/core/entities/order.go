package entities

import (
	"time"

	"github.com/sousair/americastech-exchange/internal/core/enums"
)

type (
	Order struct {
		ID         string               `json:"id"`
		ExternalID string               `json:"external_id"`
		UserID     string               `json:"user_id"`
		Pair       string               `json:"pair"`
		Direction  enums.OrderDirection `json:"direction"`
		Amount     string               `json:"amount"`
		Type       enums.OrderType      `json:"type"`
		Price      string               `json:"price"`
		Status     enums.OrderStatus    `json:"status"`
		CreatedAt  *time.Time           `json:"created_at"`
		UpdatedAt  *time.Time           `json:"updated_at"`
	}
)
