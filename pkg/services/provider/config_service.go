package services_provider

import (
	services_config "backend/pkg/services/config"
	services_interface_config "backend/pkg/services/config/interface"
)

func (object *ProviderService) ConfigService() services_interface_config.ConfigService {
	if object.configService == nil {
		object.configService = services_config.NewConfigService("config/config.yml")
	}

	return object.configService
}
