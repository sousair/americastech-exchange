package http_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type (
	GetOrderRequest struct {
		OrderId string `param:"order_id"`
	}

	GetOrderResponse struct {
		Order *entities.Order `json:"orders"`
	}

	GetOrderHandler struct {
		getOrderUseCase usecases.GetOrderUseCase
	}
)

func NewGetOrderHandler(getOrderUseCase usecases.GetOrderUseCase) *GetOrderHandler {
	return &GetOrderHandler{
		getOrderUseCase: getOrderUseCase,
	}
}

func (h GetOrderHandler) Handle(e echo.Context) error {
	userId := e.Get("user_id").(string)

	var gerOrderReq GetOrderRequest

	if err := e.Bind(&gerOrderReq); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	order, err := h.getOrderUseCase.Get(usecases.GetOrderParams{
		OrderId: gerOrderReq.OrderId,
		UserId:  userId,
	})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if order == nil {
		return e.JSON(http.StatusNotFound, map[string]string{
			"message": "order not found",
		})
	}

	return e.JSON(http.StatusOK, GetOrderResponse{
		Order: order,
	})
}
