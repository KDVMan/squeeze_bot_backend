package services_provider

import (
	services_symbol_list "backend/internal/services/symbol_list"
	services_interface_symbol_list "backend/internal/services/symbol_list/interface"
)

func (object *ProviderService) SymbolListService() services_interface_symbol_list.SymbolListService {
	if object.symbolListService == nil {
		object.symbolListService = services_symbol_list.NewSymbolListService(
			object.StorageService,
		)
	}

	return object.symbolListService
}
