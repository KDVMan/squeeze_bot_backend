package services_exchange_websocket

import (
	"backend/internal/enums"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *exchangeWebsocketServiceImplementation) SubscribeSymbol(symbol string) {
	object.symbolMutex.Lock()

	if _, exists := object.symbolsSubscriptions[symbol]; exists {
		object.symbolMutex.Unlock()
		return
	}

	stopChannel := make(chan struct{})
	object.symbolsSubscriptions[symbol] = stopChannel
	object.symbolMutex.Unlock()

	go func() {
		handler := func(event *futures.WsAggTradeEvent) {
			object.quoteRepositoryService().UpdateQuote(symbol, enums.Interval1m, event)
			object.botService().GetDealChannel() <- symbol
		}

		errorHandler := func(err error) {
			object.loggerService().Error().Printf("websocket error for trade %s: %v", symbol, err)
		}

		_, stop, err := futures.WsAggTradeServe(symbol, handler, errorHandler)
		if err != nil {
			object.loggerService().Error().Printf("failed to subscribe to trade %s: %v", symbol, err)
			return
		}

		select {
		case <-stopChannel:
			stop <- struct{}{}
		}
	}()
}

func (object *exchangeWebsocketServiceImplementation) UnsubscribeSymbol(symbol string) {
	object.symbolMutex.Lock()
	defer object.symbolMutex.Unlock()

	if stopChannel, exists := object.symbolsSubscriptions[symbol]; exists {
		close(stopChannel)
		delete(object.symbolsSubscriptions, symbol)
	} else {
		object.loggerService().Info().Printf("no active subscription found for %s\n", symbol)
	}
}
