package http_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type (
	CancelOrderRequest struct {
		OrderID string `param:"order_id"`
	}

	CancelOrderHandler struct {
		cancelOrderUseCase usecases.CancelOrderUseCase
	}
)

func NewCancelOrderHandler(cancelOrderUseCase usecases.CancelOrderUseCase) *CancelOrderHandler {
	return &CancelOrderHandler{
		cancelOrderUseCase: cancelOrderUseCase,
	}
}

func (h CancelOrderHandler) Handle(e echo.Context) error {
	var cancelOrderRequest CancelOrderRequest

	if err := e.Bind(&cancelOrderRequest); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	userId := e.Get("user_id").(string)

	err := h.cancelOrderUseCase.Cancel(usecases.CancelOrderParams{
		OrderID: cancelOrderRequest.OrderID,
		UserID:  userId,
	})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "order canceled",
	})
}
