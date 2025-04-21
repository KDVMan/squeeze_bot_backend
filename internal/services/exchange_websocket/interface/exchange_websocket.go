package services_interface_exchange_websocket

import "backend/internal/enums"

type ExchangeWebSocketService interface {
	Start()
	Stop()
	SubscribeCurrentPrice(string, enums.Interval)
	SubscribeSymbol(string)
	UnsubscribeSymbol(string)
}
