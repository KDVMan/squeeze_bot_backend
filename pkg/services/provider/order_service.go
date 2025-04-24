package services_provider

import (
	services_order "backend/internal/services/order"
	services_interface_order "backend/internal/services/order/interface"
)

func (object *ProviderService) OrderService() services_interface_order.OrderService {
	if object.orderService == nil {
		object.orderService = services_order.NewOrderService(
			object.LoggerService,
			object.StorageService,
			object.DumpService,
			object.BotRepositoryService,
			object.UserService,
		)
	}

	return object.orderService
}
