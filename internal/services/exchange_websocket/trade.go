package services_exchange_websocket

import (
	"fmt"
)

func (object *exchangeWebsocketServiceImplementation) SubscribeTrade(symbol string) {
	// object.tradeMutex.Lock()
	// defer object.tradeMutex.Unlock()
	//
	// if _, exists := object.tradesSubscriptions[symbol]; exists {
	// 	return
	// }
	//
	// stopChannel := make(chan struct{})
	// object.tradesSubscriptions[symbol] = stopChannel
	//
	// go func() {
	// 	handler := func(event *futures.WsAggTradeEvent) {
	// 		object.quoteRepositoryService().UpdateQuoteByTrade(symbol, enums.Interval1m, event)
	// 		object.tradeService().GetDealChannel() <- symbol
	// 	}
	//
	// 	errorHandler := func(err error) {
	// 		object.loggerService().Error().Printf("websocket error for trade %s: %v", symbol, err)
	// 	}
	//
	// 	_, stop, err := futures.WsAggTradeServe(symbol, handler, errorHandler)
	// 	if err != nil {
	// 		object.loggerService().Error().Printf("failed to subscribe to trade %s: %v", symbol, err)
	// 		return
	// 	}
	//
	// 	select {
	// 	case <-stopChannel:
	// 		stop <- struct{}{}
	// 	}
	// }()
}

func (object *exchangeWebsocketServiceImplementation) UnsubscribeTrade(symbol string) {
	object.tradeMutex.Lock()
	defer object.tradeMutex.Unlock()

	if stopChannel, exists := object.tradesSubscriptions[symbol]; exists {
		close(stopChannel)
		delete(object.tradesSubscriptions, symbol)
	} else {
		fmt.Printf("No active subscription found for %s\n", symbol)
	}
}
