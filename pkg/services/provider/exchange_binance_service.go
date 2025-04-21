package services_provider

import (
	services_exchange_binance "backend/pkg/services/exchange_binance"
	services_interface_exchange_binance "backend/pkg/services/exchange_binance/interface"
)

func (object *ProviderService) ExchangeBinanceService() services_interface_exchange_binance.ExchangeBinanceService {
	if object.exchangeBinanceService == nil {
		object.exchangeBinanceService = services_exchange_binance.NewExchangeBinanceService()
	}

	return object.exchangeBinanceService
}
