package services_provider

import (
	routes_symbol "backend/internal/routes/symbol"
	routes_symbol_interface "backend/internal/routes/symbol/interface"
)

func (object *ProviderService) SymbolRoute() routes_symbol_interface.SymbolRoute {
	if object.symbolRoute == nil {
		object.symbolRoute = routes_symbol.NewSymbolRoute(
			object.LoggerService,
			object.RequestService,
			object.SymbolService,
			object.ExchangeService,
		)
	}

	return object.symbolRoute
}
