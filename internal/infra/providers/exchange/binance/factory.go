package binance_exchange

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	"github.com/sousair/americastech-exchange/internal/core/enums"
)

type BinanceExchangeProvider struct {
	client               *binance.Client
	UpdateOrderEventChan chan exchange.UpdateOrderEvent
}

func NewBinanceExchangeProvider(apiKey, secret string) *BinanceExchangeProvider {
	binanceClient := binance.NewClient(apiKey, secret)
	binanceClient.NewSetServerTimeService().Do(context.Background())

	listenKey, err := binanceClient.NewStartUserStreamService().Do(context.Background())

	if err != nil {
		fmt.Println("Error starting websocket:", err)
	}

	binanceExchangeProvider := &BinanceExchangeProvider{
		client:               binanceClient,
		UpdateOrderEventChan: make(chan exchange.UpdateOrderEvent),
	}

	go func() {
		binanceExchangeProvider.startUserDataWs(listenKey)
	}()

	return binanceExchangeProvider
}

func calculateTotalPrice(order *binance.CreateOrderResponse) {
	accPrice := 0.0
	for _, fill := range order.Fills {
		price, _ := strconv.ParseFloat(fill.Price, 64)
		quantity, _ := strconv.ParseFloat(fill.Quantity, 64)
		accPrice += price * quantity
	}

	order.Price = fmt.Sprintf("%.8f", accPrice)
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
