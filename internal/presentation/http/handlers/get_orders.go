package http_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type (
	GetOrdersResponse struct {
		Orders []*entities.Order `json:"orders"`
	}

	GetOrdersHandler struct {
		getOrdersUseCase usecases.GetOrdersUseCase
	}
)

func NewGetOrdersHandler(getOrdersUseCase usecases.GetOrdersUseCase) *GetOrdersHandler {
	return &GetOrdersHandler{
		getOrdersUseCase: getOrdersUseCase,
	}
}

func (h GetOrdersHandler) Handle(e echo.Context) error {
	userId := e.Get("user_id").(string)

	orders, err := h.getOrdersUseCase.Get(usecases.GetOrderParams{
		UserID: userId,
	})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if len(orders) == 0 {
		return e.JSON(http.StatusNotFound, map[string]string{
			"message": "no orders found",
		})
	}

	return e.JSON(http.StatusOK, GetOrdersResponse{
		Orders: orders,
	})
}
