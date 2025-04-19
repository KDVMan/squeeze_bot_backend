package services_provider

import (
	routes_symbol_list "backend/internal/routes/symbol_list"
	routes_interface_symbol_list "backend/internal/routes/symbol_list/interface"
)

func (object *ProviderService) SymbolListRoute() routes_interface_symbol_list.SymbolListRoute {
	if object.symbolListRoute == nil {
		object.symbolListRoute = routes_symbol_list.NewSymbolListRoute(
			object.LoggerService,
			object.RequestService,
			object.SymbolListService,
		)
	}

	return object.symbolListRoute
}
