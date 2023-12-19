package gorm_models

import (
	"time"

	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         string               `gorm:"type:uuid;primary_key"`
	ExternalID string               `gorm:"unique_index"`
	UserID     string               `gorm:"type:uuid;not null;index"`
	Pair       string               `gorm:"not null"`
	Direction  enums.OrderDirection `gorm:"not null"`
	Amount     string               `gorm:"not null"`
	Type       enums.OrderType      `gorm:"not null"`
	Price      string               `gorm:"not null"`
	Status     enums.OrderStatus    `gorm:"not null"`
	CreatedAt  *time.Time           `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time           `gorm:"autoUpdateTime"`
}

func (o Order) ToEntity() *entities.Order {
	return &entities.Order{
		ID:         o.ID,
		ExternalID: o.ExternalID,
		UserID:     o.UserID,
		Pair:       o.Pair,
		Direction:  o.Direction,
		Amount:     o.Amount,
		Type:       o.Type,
		Price:      o.Price,
		Status:     o.Status,
		CreatedAt:  o.CreatedAt,
		UpdatedAt:  o.UpdatedAt,
	}
}
