package gorm_models

import (
	"time"

	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
	ExternalID string `gorm:"unique_index"`
	UserID     string
	Pair       string
	Direction  string
	Amount     string
	Type       string
	Price      string
	Status     string
	CreatedAt  *time.Time `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime"`
}

func (o Order) ToEntity() *entities.Order {
	return &entities.Order{
		ID:         o.ID,
		ExternalID: o.ExternalID,
		UserID:     o.UserID,
		Pair:       o.Pair,
		Direction:  enums.OrderDirection(o.Direction),
		Amount:     o.Amount,
		Type:       enums.OrderType(o.Type),
		Price:      o.Price,
		Status:     enums.OrderStatus(o.Status),
		CreatedAt:  o.CreatedAt,
		UpdatedAt:  o.UpdatedAt,
	}
}
