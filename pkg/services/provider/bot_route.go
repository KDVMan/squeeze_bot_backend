package services_provider

import (
	routes_bot "backend/internal/routes/bot"
	routes_interface_bot "backend/internal/routes/bot/interface"
)

func (object *ProviderService) BotRoute() routes_interface_bot.BotRoute {
	if object.botRoute == nil {
		object.botRoute = routes_bot.NewBotRoute(
			object.LoggerService,
			object.RequestService,
			object.BotService,
		)
	}

	return object.botRoute
}
