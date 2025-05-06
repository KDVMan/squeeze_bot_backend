package services_exchange_websocket

import (
	"backend/internal/enums"
	services_helper "backend/pkg/services/helper"
	"github.com/adshao/go-binance/v2/futures"
	"time"
)

func (object *exchangeWebsocketServiceImplementation) SubscribeSymbol(symbol string, isGuard bool) {
	object.symbolMutex.Lock()

	if _, exists := object.symbolsSubscriptions[symbol]; exists {
		object.symbolMutex.Unlock()
		return
	}

	stopChannel := make(chan struct{})
	object.symbolsSubscriptions[symbol] = stopChannel
	object.symbolMutex.Unlock()

	go func() {
		for {
			doneChannel := make(chan struct{})

			handler := func(event *futures.WsAggTradeEvent) {
				object.quoteRepositoryService().UpdateQuote(symbol, enums.Interval1m, event)

				if isGuard {
					object.botService().GetGuardChannel() <- symbol
				} else {
					object.botService().GetDealChannel() <- symbol
				}
			}

			errorHandler := func(err error) {
				object.loggerService().Error().Printf("websocket error for %s: %v", symbol, err)

				select {
				case <-doneChannel:
				default:
					object.loggerService().Info().Printf("reconnect: %s", symbol)
					close(doneChannel)
				}
			}

			_, stop, err := futures.WsAggTradeServe(symbol, handler, errorHandler)
			if err != nil {
				object.loggerService().Error().Printf("failed to subscribe to %s: %v", symbol, err)
				<-time.After(object.reconnectDelay)
				continue
			}

			object.loggerService().Info().Printf("started: %s", symbol)

			select {
			case <-stopChannel:
				object.loggerService().Info().Printf("shutdown: %s", symbol)
				services_helper.SafeSendSignal(stop)
				return
			case <-object.doneChannel:
				object.loggerService().Info().Printf("shutdown (system): %s", symbol)
				services_helper.SafeSendSignal(stop)
				return
			case <-doneChannel:
				object.loggerService().Info().Printf("reconnecting... %s", symbol)
				services_helper.SafeSendSignal(stop)
				<-time.After(object.reconnectDelay)
				continue
			}
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
