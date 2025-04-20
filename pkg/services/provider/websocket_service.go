package services_provider

import (
	services_websocket "backend/internal/services/websocket"
	services_interface_websocket "backend/internal/services/websocket/interface"
)

func (object *ProviderService) WebsocketService() services_interface_websocket.WebsocketService {
	if object.websocketService == nil {
		object.websocketService = services_websocket.NewWebsocketService(
			object.LoggerService,
			object.ExchangeLimitService,
			object.SymbolService,
			object.UserService,
			object.BotService,
		)
	}

	return object.websocketService
}
