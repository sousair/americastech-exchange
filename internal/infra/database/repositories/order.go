package gorm_repositories

import (
	"github.com/google/uuid"
	"github.com/sousair/americastech-exchange/internal/application/providers/repositories"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	gorm_models "github.com/sousair/americastech-exchange/internal/infra/database/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repositories.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(order *entities.Order) (*entities.Order, error) {
	gormOrder := gorm_models.Order{
		ID:         uuid.New().String(),
		ExternalID: order.ExternalID,
		UserID:     order.UserID,
		Pair:       order.Pair,
		Direction:  string(order.Direction),
		Amount:     order.Amount,
		Type:       string(order.Type),
		Price:      order.Price,
		Status:     string(order.Status),
	}

	if err := r.db.Create(&gormOrder).Error; err != nil {
		return nil, err
	}

	return gormOrder.ToEntity(), nil
}
