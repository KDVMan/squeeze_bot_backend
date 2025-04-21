package services_provider

import (
	services_exchange_websocket "backend/internal/services/exchange_websocket"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
)

func (object *ProviderService) ExchangeWebsocketService() services_interface_exchange_websocket.ExchangeWebSocketService {
	if object.exchangeWebsocketService == nil {
		object.exchangeWebsocketService = services_exchange_websocket.NewExchangeWebsocketService(
			object.LoggerService,
			object.SymbolService,
			object.QuoteService,
			object.QuoteRepositoryService,
			object.ExchangeService,
			object.UserService,
			object.BotService,
		)
	}

	return object.exchangeWebsocketService
}
