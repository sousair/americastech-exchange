package http_handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-exchange/internal/application/errors"
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

	if userId == "" {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": "user_id not found",
		})
	}

	orders, err := h.getOrdersUseCase.Get(usecases.GetOrdersParams{
		UserID: userId,
	})

	if err != nil {
		if errors.As(err, &custom_errors.OrderNotFoundError) {
			return e.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, GetOrdersResponse{
		Orders: orders,
	})
}
