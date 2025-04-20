package services_init

import (
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_init "backend/internal/services/init/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type initServiceImplementation struct {
	storageService   func() services_interface_storage.StorageService
	websocketService func() services_interface_websocket.WebsocketService
	symbolService    func() services_interface_symbol.SymbolService
	botService       func() services_interface_bot.BotService
}

func NewInitService(
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	symbolService func() services_interface_symbol.SymbolService,
	botService func() services_interface_bot.BotService,
) services_interface_init.InitService {
	return &initServiceImplementation{
		storageService:   storageService,
		websocketService: websocketService,
		symbolService:    symbolService,
		botService:       botService,
	}
}
