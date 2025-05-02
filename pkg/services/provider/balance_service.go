package services_provider

import (
	services_balance "backend/internal/services/balance"
	services_interface_balance "backend/internal/services/balance/interface"
)

func (object *ProviderService) BalanceService() services_interface_balance.BalanceService {
	if object.balanceService == nil {
		object.balanceService = services_balance.NewBalanceService(
			object.LoggerService,
		)
	}

	return object.balanceService
}
