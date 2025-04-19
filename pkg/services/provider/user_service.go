package services_provider

import (
	services_user "backend/internal/services/user"
	services_interface_user "backend/internal/services/user/interface"
)

func (object *ProviderService) UserService() services_interface_user.UserService {
	if object.userService == nil {
		object.userService = services_user.NewUserService(
			object.StorageService,
			object.WebsocketService,
			object.ExchangeService,
		)
	}

	return object.userService
}
