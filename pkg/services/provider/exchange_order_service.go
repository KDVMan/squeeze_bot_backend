package services_provider

import (
	services_exchange_order "backend/internal/services/exchange_order"
	services_interface_exchange_order "backend/internal/services/exchange_order/interface"
)

func (object *ProviderService) ExchangeOrderService() services_interface_exchange_order.ExchangeOrderService {
	if object.exchangeOrderService == nil {
		object.exchangeOrderService = services_exchange_order.NewExchangeOrderService(
			object.LoggerService,
			object.ConfigService,
			object.DumpService,
		)
	}

	return object.exchangeOrderService
}
