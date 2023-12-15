package binance_exchange

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	"github.com/sousair/americastech-exchange/internal/core/entities"
	"github.com/sousair/americastech-exchange/internal/core/enums"
)

func (b BinanceExchangeProvider) Create(params exchange.CreateOrderParams) (*exchange.CreatedOrder, error) {
	var binanceOrderService *binance.CreateOrderService

	orderService := b.client.NewCreateOrderService().
		Symbol(params.Pair).
		Side(binance.SideType(params.Direction)).
		Quantity(params.Amount)

	switch params.Type {
	case enums.OrderTypeMarket:
		binanceOrderService = orderService.Type("MARKET")

	case enums.OrderTypeLimit:
		binanceOrderService = orderService.Type("LIMIT").
			TimeInForce(binance.TimeInForceTypeGTC).
			Price(params.Price)

	}

	binanceOrderRes, err := binanceOrderService.Do(context.Background())

	if err != nil {
		return nil, err
	}

	if binanceOrderRes.Status == binance.OrderStatusTypeFilled {
		totalQuantity := float64(0)
		totalPrice := float64(0)

		for _, fill := range binanceOrderRes.Fills {
			quantity, _ := strconv.ParseFloat(fill.Quantity, 64)
			price, _ := strconv.ParseFloat(fill.Price, 64)

			totalQuantity += quantity
			totalPrice += quantity * price
		}

		binanceOrderRes.Price = fmt.Sprintf("%.8f", totalPrice/totalQuantity)
	}

	createdOrder := &exchange.CreatedOrder{
		ExternalID: fmt.Sprint(binanceOrderRes.OrderID),
		Pair:       binanceOrderRes.Symbol,
		Direction:  enums.OrderDirection(binanceOrderRes.Side),
		Type:       enums.OrderType(binanceOrderRes.Type),
		Amount:     binanceOrderRes.OrigQuantity,
		Price:      binanceOrderRes.Price,
		Status:     binanceStatusMap[binanceOrderRes.Status],
	}

	return createdOrder, nil
}

func (b BinanceExchangeProvider) CancelOrder(order *entities.Order) error {
	orderID, err := strconv.ParseInt(order.ExternalID, 10, 64)
	if err != nil {
		return err
	}

	_, err = b.client.NewCancelOrderService().
		Symbol(order.Pair).
		OrderID(orderID).
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
