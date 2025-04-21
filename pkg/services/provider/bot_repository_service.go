package services_provider

import (
	services_bot_repository "backend/internal/services/bot_repository"
	services_interface_bot_repository "backend/internal/services/bot_repository/interface"
)

func (object *ProviderService) BotRepositoryService() services_interface_bot_repository.BotRepositoryService {
	if object.botRepositoryService == nil {
		object.botRepositoryService = services_bot_repository.NewBotRepositoryService(
			object.LoggerService,
		)
	}

	return object.botRepositoryService
}
