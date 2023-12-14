package binance_exchange

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	"github.com/sousair/americastech-exchange/internal/core/enums"
)

type (
	BinanceExchangeProvider struct {
		client *binance.Client
	}
)

func NewBinanceExchangeProvider(apiKey, secret, baseUrl string) exchange.ExchangeProvider {
	binanceClient := binance.NewClient(apiKey, secret)
	binanceClient.BaseURL = baseUrl
	binanceClient.TimeOffset = -3000

	return &BinanceExchangeProvider{
		client: binanceClient,
	}
}

func (b BinanceExchangeProvider) Create(params exchange.CreateOrderParams) (*exchange.CreatedOrder, error) {
	var binanceOrderService *binance.CreateOrderService

	switch params.Type {
	case "MARKET":
		binanceOrderService = b.mountMarketOrder(params.Pair, string(params.Direction), params.Amount)
	case "LIMIT":
		binanceOrderService = b.mountLimitOrder(params.Pair, string(params.Direction), params.Amount, params.Price)
	}

	binanceOrderRes, err := binanceOrderService.Do(context.Background())

	if err != nil {
		return nil, err
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

func parseBinanceStatus(status binance.OrderStatusType) enums.OrderStatus {
	switch status {
	case binance.OrderStatusTypeNew:
		return enums.Open
	case binance.OrderStatusTypePartiallyFilled:
		return enums.PartiallyFilled
	case binance.OrderStatusTypeFilled:
		return enums.Filled
	case binance.OrderStatusTypeCanceled:
		return enums.Canceled
	default:
		return ""
	}
}
