package binance_exchange

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/sousair/americastech-exchange/internal/application/providers/exchange"
	"github.com/sousair/americastech-exchange/internal/core/enums"
)

func (b BinanceExchangeProvider) startUserDataWs(listenKey string) {
	doneC, _, err := binance.WsUserDataServe(listenKey, b.wsUserDataHandler, b.wsUserDataErrHandler)

	if err != nil {
		fmt.Println("Error serving websocket:", err)
	}

	fmt.Println("Websocket started")
	<-doneC
}

func (b BinanceExchangeProvider) wsUserDataHandler(event *binance.WsUserDataEvent) {
	fmt.Printf("Event: %v\n", event)
	if event.Event == binance.UserDataEventTypeExecutionReport {
		orderEvent := event.OrderUpdate
		b.UpdateOrderEventChan <- exchange.UpdateOrderEvent{
			ExternalID: fmt.Sprint(orderEvent.Id),
			Pair:       orderEvent.Symbol,
			Direction:  enums.OrderDirection(orderEvent.Side),
			Type:       enums.OrderType(orderEvent.Type),
			Price:      orderEvent.Price,
			Status:     parseBinanceStatus(binance.OrderStatusType(orderEvent.Status)),
		}
	}

}

func (b BinanceExchangeProvider) wsUserDataErrHandler(err error) {
	fmt.Println(err)
}
