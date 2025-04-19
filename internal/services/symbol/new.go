package services_symbol

import (
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type symbolServiceImplementation struct {
	loggerService    func() services_interface_logger.LoggerService
	storageService   func() services_interface_storage.StorageService
	websocketService func() services_interface_websocket.WebsocketService
}

func NewSymbolService(
	loggerService func() services_interface_logger.LoggerService,
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
) services_interface_symbol.SymbolService {
	return &symbolServiceImplementation{
		loggerService:    loggerService,
		storageService:   storageService,
		websocketService: websocketService,
	}
}
