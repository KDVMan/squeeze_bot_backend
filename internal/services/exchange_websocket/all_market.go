package services_exchange_websocket

import (
	"github.com/adshao/go-binance/v2/futures"
	"time"
)

func (object *exchangeWebsocketServiceImplementation) allMarket() {
	for {
		var stopChannel chan struct{}

		doneChannel, stopChannel, err := futures.WsAllMarketTickerServe(
			func(event futures.WsAllMarketTickerEvent) {
				if err := object.symbolService().UpdateStatistic(event); err != nil {
					object.loggerService().Error().Printf("failed to update statistic: %v", err)
				}
			},
			func(err error) {
				object.loggerService().Error().Printf("websocket error: %v", err)
				stopChannel <- struct{}{} // корректно закрываем текущее соединение
			},
		)
		if err != nil {
			object.loggerService().Error().Printf("failed to start websocket: %v", err)
			time.Sleep(object.reconnectDelay)
			continue
		}

		object.loggerService().Info().Println("all market - started")

		select {
		case <-doneChannel:
			object.loggerService().Info().Printf("reconnect")
			time.Sleep(object.reconnectDelay)
			continue
		case <-object.doneChannel:
			stopChannel <- struct{}{}
			return
		}
	}
}
