package http_handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-exchange/internal/application/errors"
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

	if userId == "" {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": "user_id not found",
		})
	}

	err := h.cancelOrderUseCase.Cancel(usecases.CancelOrderParams{
		OrderID: cancelOrderRequest.OrderID,
		UserID:  userId,
	})

	if err != nil {
		if errors.As(err, &custom_errors.OrderNotFoundError) {
			return e.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		}

		if errors.As(err, &custom_errors.OrderAlreadyCanceledError) {
			return e.JSON(http.StatusConflict, map[string]string{
				"message": err.Error(),
			})
		}

		if errors.As(err, &custom_errors.OrderAlreadyFilledError) {
			return e.JSON(http.StatusConflict, map[string]string{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "order canceled",
	})
}
