package services_exchange_limit

import (
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type exchangeLimitServiceImplementation struct {
	storageService   func() services_interface_storage.StorageService
	websocketService func() services_interface_websocket.WebsocketService
}

func NewExchangeLimitService(
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
) services_interface_exchange_limit.ExchangeLimitService {
	return &exchangeLimitServiceImplementation{
		storageService:   storageService,
		websocketService: websocketService,
	}
}
