package http_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type (
	CreateOrderRequest struct {
		Pair      string `json:"pair"`
		Direction string `json:"direction"`
		Type      string `json:"type"`
		Amount    string `json:"amount"`
		Price     string `json:"price"`
	}

	CreateOrderResponse struct {
		Order *entities.Order `json:"order"`
	}

	CreateOrderHandler struct {
		createOrderUseCase usecases.CreateOrderUseCase
	}
)

func NewCreateOrderHandler(createOrderUseCase usecases.CreateOrderUseCase) *CreateOrderHandler {
	return &CreateOrderHandler{
		createOrderUseCase: createOrderUseCase,
	}
}

func (h CreateOrderHandler) Handle(e echo.Context) error {
	var request CreateOrderRequest

	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	userId := e.Get("user_id").(string)

	order, err := h.createOrderUseCase.Create(usecases.CreateOrderParams{
		UserID:    userId,
		Pair:      request.Pair,
		Direction: enums.OrderDirection(request.Direction),
		Type:      enums.OrderType(request.Type),
		Amount:    request.Amount,
		Price:     request.Price,
	})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
	}

	return e.JSON(http.StatusCreated, CreateOrderResponse{
		Order: order,
	})
}
