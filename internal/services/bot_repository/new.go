package services_bot_repository

import (
	models_bot "backend/internal/models/bot"
	services_interface_bot_repository "backend/internal/services/bot_repository/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
)

type botRepositoryServiceImplementation struct {
	loggerService func() services_interface_logger.LoggerService
	data          map[string][]*models_bot.BotModel
	mutex         *sync.Mutex
}

func NewBotRepositoryService(
	loggerService func() services_interface_logger.LoggerService,
) services_interface_bot_repository.BotRepositoryService {
	return &botRepositoryServiceImplementation{
		loggerService: loggerService,
		data:          make(map[string][]*models_bot.BotModel),
		mutex:         &sync.Mutex{},
	}
}
