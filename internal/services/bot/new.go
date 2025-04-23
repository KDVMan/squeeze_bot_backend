package services_bot

import (
	models_bot "backend/internal/models/bot"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_bot_repository "backend/internal/services/bot_repository/interface"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_exchange_order "backend/internal/services/exchange_order/interface"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
	services_interface_init "backend/internal/services/init/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_quote_repository "backend/internal/services/quote_repository/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type botServiceImplementation struct {
	loggerService            func() services_interface_logger.LoggerService
	configService            func() services_interface_config.ConfigService
	storageService           func() services_interface_storage.StorageService
	websocketService         func() services_interface_websocket.WebsocketService
	dumpService              func() services_interface_dump.DumpService
	exchangeService          func() services_interface_exchange.ExchangeService
	exchangeWebsocketService func() services_interface_exchange_websocket.ExchangeWebSocketService
	exchangeOrderService     func() services_interface_exchange_order.ExchangeOrderService
	initService              func() services_interface_init.InitService
	symbolService            func() services_interface_symbol.SymbolService
	quoteService             func() services_interface_quote.QuoteService
	quoteRepositoryService   func() services_interface_quote_repository.QuoteRepositoryService
	botRepositoryService     func() services_interface_bot_repository.BotRepositoryService
	runChannel               chan *models_bot.BotModel
	dealChannel              chan string
	addDealChannel           chan *models_bot.BotModel
	commission               float64
}

func NewBotService(
	loggerService func() services_interface_logger.LoggerService,
	configService func() services_interface_config.ConfigService,
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	dumpService func() services_interface_dump.DumpService,
	exchangeService func() services_interface_exchange.ExchangeService,
	exchangeWebsocketService func() services_interface_exchange_websocket.ExchangeWebSocketService,
	exchangeOrderService func() services_interface_exchange_order.ExchangeOrderService,
	initService func() services_interface_init.InitService,
	symbolService func() services_interface_symbol.SymbolService,
	quoteService func() services_interface_quote.QuoteService,
	quoteRepositoryService func() services_interface_quote_repository.QuoteRepositoryService,
	botRepositoryService func() services_interface_bot_repository.BotRepositoryService,
) services_interface_bot.BotService {
	return &botServiceImplementation{
		loggerService:            loggerService,
		configService:            configService,
		storageService:           storageService,
		websocketService:         websocketService,
		dumpService:              dumpService,
		exchangeService:          exchangeService,
		exchangeWebsocketService: exchangeWebsocketService,
		exchangeOrderService:     exchangeOrderService,
		initService:              initService,
		symbolService:            symbolService,
		quoteService:             quoteService,
		quoteRepositoryService:   quoteRepositoryService,
		botRepositoryService:     botRepositoryService,
		runChannel:               make(chan *models_bot.BotModel, 1000000),
		dealChannel:              make(chan string, 1000000),
		addDealChannel:           make(chan *models_bot.BotModel, 1000000),
		commission:               configService().GetConfig().Binance.FuturesCommission,
	}
}
