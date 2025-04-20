package services_provider

import (
	services_bot "backend/internal/services/bot"
	services_interface_bot "backend/internal/services/bot/interface"
)

func (object *ProviderService) BotService() services_interface_bot.BotService {
	if object.botService == nil {
		object.botService = services_bot.NewBotService(
			object.LoggerService,
			object.StorageService,
			object.WebsocketService,
			object.InitService,
			object.SymbolService,
		)
	}

	return object.botService
}
