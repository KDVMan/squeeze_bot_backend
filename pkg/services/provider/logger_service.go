package services_provider

import (
	services_logger "backend/pkg/services/logger"
	services_interface_logger "backend/pkg/services/logger/interface"
)

func (object *ProviderService) LoggerService() services_interface_logger.LoggerService {
	if object.loggerService == nil {
		object.loggerService = services_logger.NewLoggerService(
			object.ConfigService,
		)
	}

	return object.loggerService
}
