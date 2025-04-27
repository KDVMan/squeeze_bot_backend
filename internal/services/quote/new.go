package services_quote

import (
	services_interface_bot_repository "backend/internal/services/bot_repository/interface"
	services_exchange_interface "backend/internal/services/exchange/interface"
	services_exchange_websocket_interface "backend/internal/services/exchange_websocket/interface"
	services_interface_init "backend/internal/services/init/interface"
	services_quote_interface "backend/internal/services/quote/interface"
	services_websocket_interface "backend/internal/services/websocket/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_storage_interface "backend/pkg/services/storage/interface"
)

type quoteServiceImplementation struct {
	loggerService            func() services_interface_logger.LoggerService
	storageService           func() services_storage_interface.StorageService
	websocketService         func() services_websocket_interface.WebsocketService
	exchangeService          func() services_exchange_interface.ExchangeService
	exchangeWebsocketService func() services_exchange_websocket_interface.ExchangeWebSocketService
	initService              func() services_interface_init.InitService
	botRepositoryService     func() services_interface_bot_repository.BotRepositoryService
}

func NewQuoteService(
	loggerService func() services_interface_logger.LoggerService,
	storageService func() services_storage_interface.StorageService,
	websocketService func() services_websocket_interface.WebsocketService,
	exchangeService func() services_exchange_interface.ExchangeService,
	exchangeWebsocketService func() services_exchange_websocket_interface.ExchangeWebSocketService,
	initService func() services_interface_init.InitService,
	botRepositoryService func() services_interface_bot_repository.BotRepositoryService,
) services_quote_interface.QuoteService {
	return &quoteServiceImplementation{
		loggerService:            loggerService,
		storageService:           storageService,
		websocketService:         websocketService,
		exchangeService:          exchangeService,
		exchangeWebsocketService: exchangeWebsocketService,
		initService:              initService,
		botRepositoryService:     botRepositoryService,
	}
}
