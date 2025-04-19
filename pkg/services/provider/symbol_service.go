package services_provider

import (
	services_symbol "backend/internal/services/symbol"
	services_interface_symbol "backend/internal/services/symbol/interface"
)

func (object *ProviderService) SymbolService() services_interface_symbol.SymbolService {
	if object.symbolService == nil {
		object.symbolService = services_symbol.NewSymbolService(
			object.LoggerService,
			object.StorageService,
			object.WebsocketService,
		)
	}

	return object.symbolService
}
