package services_provider

import (
	services_router "backend/pkg/services/router"
	services_interface_router "backend/pkg/services/router/interface"
)

func (object *ProviderService) RouterService() services_interface_router.RouterService {
	if object.routerService == nil {
		object.routerService = services_router.NewRouterService(
			object.LoggerService,
			object.WebsocketService,
			object.InitRoute,
			object.SymbolRoute,
			object.SymbolListRoute,
			object.QuoteRoute,
			object.ChartSettingsRoute,
			object.BotRoute,
		)
	}

	return object.routerService
}
