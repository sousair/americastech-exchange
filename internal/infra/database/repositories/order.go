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

func (r OrderRepository) Create(order *entities.Order) (*entities.Order, error) {
	gormOrder := gorm_models.Order{
		ID:         uuid.New().String(),
		ExternalID: order.ExternalID,
		UserID:     order.UserID,
		Pair:       order.Pair,
		Direction:  order.Direction,
		Amount:     order.Amount,
		Type:       order.Type,
		Price:      order.Price,
		Status:     order.Status,
	}

	if err := r.db.Create(&gormOrder).Error; err != nil {
		return nil, err
	}

	return gormOrder.ToEntity(), nil
}

func (r OrderRepository) FindAllBy(params map[string]interface{}) ([]*entities.Order, error) {
	var gormOrders []*gorm_models.Order

	if err := r.db.Where(params).Find(&gormOrders).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	var orders []*entities.Order

	for _, gormOrder := range gormOrders {
		orders = append(orders, gormOrder.ToEntity())
	}

	return orders, nil
}

func (r OrderRepository) FindOneBy(params map[string]interface{}) (*entities.Order, error) {
	var gormOrder gorm_models.Order

	if err := r.db.Where(params).First(&gormOrder).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return gormOrder.ToEntity(), nil
}

func (r OrderRepository) Update(order *entities.Order) (*entities.Order, error) {
	orderModel := gorm_models.Order{
		ID:         order.ID,
		ExternalID: order.ExternalID,
		UserID:     order.UserID,
		Pair:       order.Pair,
		Direction:  order.Direction,
		Amount:     order.Amount,
		Type:       order.Type,
		Price:      order.Price,
		Status:     order.Status,
	}

	if err := r.db.Model(orderModel).Updates(&orderModel).Error; err != nil {
		return nil, err
	}

	return r.FindOneBy(map[string]interface{}{
		"id": order.ID,
	})
}
