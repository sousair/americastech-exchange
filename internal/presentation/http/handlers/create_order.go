package http_handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/enums"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
)

type (
	CreateOrderRequest struct {
		Pair      string `json:"pair" validate:"required"`
		Direction string `json:"direction" validate:"required,oneof=BUY SELL"`
		Type      string `json:"type" validate:"required,oneof=MARKET LIMIT"`
		Amount    string `json:"amount" validate:"required,numeric"`
		Price     string `json:"price" validate:"required_if=Type LIMIT,omitempty,numeric"`
	}

	CreateOrderResponse struct {
		Order *entities.Order `json:"order"`
	}

	CreateOrderHandler struct {
		createOrderUseCase usecases.CreateOrderUseCase
		validator          *validator.Validate
	}
)

func NewCreateOrderHandler(createOrderUseCase usecases.CreateOrderUseCase, validator *validator.Validate) *CreateOrderHandler {
	return &CreateOrderHandler{
		createOrderUseCase: createOrderUseCase,
		validator:          validator,
	}
}

func (h CreateOrderHandler) Handle(e echo.Context) error {
	var request CreateOrderRequest

	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	if err := h.validator.Struct(request); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
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
