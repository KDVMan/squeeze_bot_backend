package services_provider

import (
	services_init "backend/internal/services/init"
	services_interface_init "backend/internal/services/init/interface"
)

func (object *ProviderService) InitService() services_interface_init.InitService {
	if object.initService == nil {
		object.initService = services_init.NewInitService(
			object.StorageService,
			object.WebsocketService,
			object.SymbolService,
			// object.CalculatorService,
			// object.CalculatorPresetService,
		)
	}

	return object.initService
}
