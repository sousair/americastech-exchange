package repositories

import "github.com/sousair/americastech-exchange/internal/core/entities"

type (
	OrderRepository interface {
		Create(order *entities.Order) (*entities.Order, error)
		FindOneBy(where map[string]interface{}) (*entities.Order, error)
		FindAllBy(where map[string]interface{}) ([]*entities.Order, error)
		Update(params *entities.Order) (*entities.Order, error)
	}
)
