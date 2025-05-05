package services_exchange_websocket

import (
	"backend/internal/enums"
	services_helper "backend/pkg/services/helper"
	"github.com/adshao/go-binance/v2/futures"
	"time"
)

func (object *exchangeWebsocketServiceImplementation) SubscribeCurrentPrice(symbol string, interval enums.Interval) {
	if object.currentPriceSymbol == symbol && object.currentPriceInterval == interval {
		return
	}

	if object.currentPriceStopChannel != nil {
		close(object.currentPriceStopChannel)
		object.currentPriceStopChannel = nil
	}

	stopChannel := make(chan struct{})
	object.currentPriceStopChannel = stopChannel
	object.currentPriceSymbol = symbol
	object.currentPriceInterval = interval

	go func() {
		for {
			doneChannel := make(chan struct{})

			object.loggerService().Info().Printf("starting...")

			handler := func(event *futures.WsKlineEvent) {
				if err := object.quoteService().CurrentPrice(symbol, interval, event.Kline); err != nil {
					object.loggerService().Error().Printf("failed to update current price: %v", err)
				}
			}

			errorHandler := func(err error) {
				object.loggerService().Error().Printf("websocket error for %s: %v", symbol, err)

				select {
				case <-doneChannel:
				default:
					object.loggerService().Info().Printf("reconnect")
					close(doneChannel)
				}
			}

			_, stop, err := futures.WsKlineServe(symbol, interval.String(), handler, errorHandler)
			if err != nil {
				object.loggerService().Error().Printf("failed to subscribe to %s: %v", symbol, err)
				time.Sleep(object.reconnectDelay)
				continue
			}

			object.loggerService().Info().Printf("started")

			select {
			case <-stopChannel:
				object.loggerService().Info().Printf("shutdown")
				services_helper.SafeSendSignal(stop)
				return
			case <-doneChannel:
				object.loggerService().Info().Printf("reconnecting...")
				services_helper.SafeSendSignal(stop)
				time.Sleep(object.reconnectDelay)
				continue
			}
		}
	}()
}
