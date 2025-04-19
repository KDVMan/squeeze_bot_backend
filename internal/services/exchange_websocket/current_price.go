package services_exchange_websocket

import (
	"backend/internal/enums"
	"github.com/adshao/go-binance/v2/futures"
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
		handler := func(event *futures.WsKlineEvent) {
			if err := object.quoteService().CurrentPrice(symbol, interval, event.Kline); err != nil {
				object.loggerService().Error().Printf("failed to update current price: %v", err)
			}
		}

		errorHandler := func(err error) {
			object.loggerService().Error().Printf("webSocket error for %s: %v", object.currentPriceSymbol, err)
		}

		_, stop, err := futures.WsKlineServe(object.currentPriceSymbol, object.currentPriceInterval.String(), handler, errorHandler)
		if err != nil {
			object.loggerService().Error().Printf("failed to subscribe to %s: %v", object.currentPriceSymbol, err)
			return
		}

		select {
		case <-stopChannel:
			stop <- struct{}{}
		}
	}()
}
