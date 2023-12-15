package binance_exchange

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	"github.com/sousair/americastech-exchange/internal/core/enums"
)

var binanceStatusMap = map[binance.OrderStatusType]enums.OrderStatus{
	binance.OrderStatusTypeNew:             enums.OrderStatusOpen,
	binance.OrderStatusTypePartiallyFilled: enums.OrderStatusPartiallyFilled,
	binance.OrderStatusTypeFilled:          enums.OrderStatusFilled,
	binance.OrderStatusTypeCanceled:        enums.OrderStatusCanceled,
}

type BinanceExchangeProvider struct {
	client               *binance.Client
	UpdateOrderEventChan chan exchange.UpdateOrderEvent
}

func NewBinanceExchangeProvider(apiKey, secret string) *BinanceExchangeProvider {
	binanceClient := binance.NewClient(apiKey, secret)
	binanceClient.NewSetServerTimeService().Do(context.Background())

	listenKey, err := binanceClient.NewStartUserStreamService().Do(context.Background())

	if err != nil {
		fmt.Println("[Binance] Error starting websocket:", err)
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
