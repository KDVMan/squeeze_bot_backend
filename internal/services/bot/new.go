package services_bot

import (
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_init "backend/internal/services/init/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type botServiceImplementation struct {
	loggerService    func() services_interface_logger.LoggerService
	storageService   func() services_interface_storage.StorageService
	websocketService func() services_interface_websocket.WebsocketService
	initService      func() services_interface_init.InitService
	symbolService    func() services_interface_symbol.SymbolService
	// dealChannel                    chan string
}

func NewBotService(
	loggerService func() services_interface_logger.LoggerService,
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	initService func() services_interface_init.InitService,
	symbolService func() services_interface_symbol.SymbolService,
) services_interface_bot.BotService {
	return &botServiceImplementation{
		loggerService:    loggerService,
		storageService:   storageService,
		websocketService: websocketService,
		initService:      initService,
		symbolService:    symbolService,
		// dealChannel:                    make(chan string, 1000000),
	}
}
