package services_quote

import (
	services_exchange_interface "backend/internal/services/exchange/interface"
	services_exchange_websocket_interface "backend/internal/services/exchange_websocket/interface"
	services_quote_interface "backend/internal/services/quote/interface"
	services_websocket_interface "backend/internal/services/websocket/interface"
	services_storage_interface "backend/pkg/services/storage/interface"
)

type quoteServiceImplementation struct {
	storageService           func() services_storage_interface.StorageService
	websocketService         func() services_websocket_interface.WebsocketService
	exchangeService          func() services_exchange_interface.ExchangeService
	exchangeWebsocketService func() services_exchange_websocket_interface.ExchangeWebSocketService
}

func NewQuoteService(
	storageService func() services_storage_interface.StorageService,
	websocketService func() services_websocket_interface.WebsocketService,
	exchangeService func() services_exchange_interface.ExchangeService,
	exchangeWebsocketService func() services_exchange_websocket_interface.ExchangeWebSocketService,
) services_quote_interface.QuoteService {
	return &quoteServiceImplementation{
		storageService:           storageService,
		websocketService:         websocketService,
		exchangeService:          exchangeService,
		exchangeWebsocketService: exchangeWebsocketService,
	}
}
