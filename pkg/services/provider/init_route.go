package services_provider

import (
	routes_init "backend/internal/routes/init"
	routes_init_interface "backend/internal/routes/init/interface"
)

func (object *ProviderService) InitRoute() routes_init_interface.InitRoute {
	if object.initRoute == nil {
		object.initRoute = routes_init.NewInitRoute(
			object.LoggerService,
			object.RequestService,
			object.InitService,
		)
	}

	return object.initRoute
}
