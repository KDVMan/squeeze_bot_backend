package services_provider

import (
	services_exchange "backend/internal/services/exchange"
	services_exchange_interface "backend/internal/services/exchange/interface"
)

func (object *ProviderService) ExchangeService() services_exchange_interface.ExchangeService {
	if object.exchangeService == nil {
		object.exchangeService = services_exchange.NewExchangeService(
			object.ConfigService,
			object.StorageService,
			object.DumpService,
			object.ExchangeLimitService,
		)
	}

	return object.exchangeService
}
