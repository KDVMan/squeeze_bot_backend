package services_exchange_websocket

import (
	services_helper "backend/pkg/services/helper"
	"github.com/adshao/go-binance/v2/futures"
	"time"
)

func (object *exchangeWebsocketServiceImplementation) allMarket() {
	for {
		doneChannel := make(chan struct{})

		object.loggerService().Info().Printf("starting...")

		handler := func(event futures.WsAllMarketTickerEvent) {
			if err := object.symbolService().UpdateStatistic(event); err != nil {
				object.loggerService().Error().Printf("failed to update statistic: %v", err)
			}
		}

		errorHandler := func(err error) {
			object.loggerService().Error().Printf("websocket error: %v", err)

			select {
			case <-doneChannel:
			default:
				object.loggerService().Info().Printf("reconnect")
				close(doneChannel)
			}
		}

		_, stopChannel, err := futures.WsAllMarketTickerServe(handler, errorHandler)
		if err != nil {
			object.loggerService().Error().Printf("failed to start websocket: %v", err)
			time.Sleep(object.reconnectDelay)
			continue
		}

		object.loggerService().Info().Printf("started")

		select {
		case <-object.doneChannel:
			object.loggerService().Info().Printf("shutdown")
			stopChannel <- struct{}{}
			return
		case <-doneChannel:
			object.loggerService().Info().Printf("reconnecting...")
			services_helper.SafeSendSignal(stopChannel)
			time.Sleep(object.reconnectDelay)
			continue
		}
	}
}
