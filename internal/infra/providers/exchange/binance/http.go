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

	switch params.Type {
	case enums.Market:
		binanceOrderService = b.mountMarketOrder(params.Pair, string(params.Direction), params.Amount)
	case enums.Limit:
		binanceOrderService = b.mountLimitOrder(params.Pair, string(params.Direction), params.Amount, params.Price)
	}

	binanceOrderRes, err := binanceOrderService.Do(context.Background())

	if err != nil {
		return nil, err
	}

	if binanceOrderRes.Status == binance.OrderStatusTypeFilled {
		calculateTotalPrice(binanceOrderRes)
	}

	createdOrder := &exchange.CreatedOrder{
		ExternalID: fmt.Sprint(binanceOrderRes.OrderID),
		Pair:       binanceOrderRes.Symbol,
		Direction:  enums.OrderDirection(binanceOrderRes.Side),
		Type:       enums.OrderType(binanceOrderRes.Type),
		Amount:     binanceOrderRes.OrigQuantity,
		Price:      binanceOrderRes.Price,
		Status:     parseBinanceStatus(binanceOrderRes.Status),
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

func (b BinanceExchangeProvider) mountMarketOrder(pair, direction, quantity string) *binance.CreateOrderService {
	return b.client.
		NewCreateOrderService().
		Symbol(pair).
		Side(binance.SideType(direction)).
		Type("MARKET").
		Quantity(quantity)
}

func (b BinanceExchangeProvider) mountLimitOrder(pair, direction, quantity, price string) *binance.CreateOrderService {
	return b.client.
		NewCreateOrderService().
		Symbol(pair).
		Side(binance.SideType(direction)).
		Type("LIMIT").
		TimeInForce(binance.TimeInForceTypeGTC).
		Quantity(quantity).
		Price(price)
}
