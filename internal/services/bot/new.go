package services_bot

import (
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type botServiceImplementation struct {
	storageService   func() services_interface_storage.StorageService
	websocketService func() services_interface_websocket.WebsocketService
	symbolService    func() services_interface_symbol.SymbolService
	// dealChannel                    chan string
}

func NewBotService(
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	symbolService func() services_interface_symbol.SymbolService,
) services_interface_bot.BotService {
	return &botServiceImplementation{
		storageService:   storageService,
		websocketService: websocketService,
		symbolService:    symbolService,
		// dealChannel:                    make(chan string, 1000000),
	}
}
