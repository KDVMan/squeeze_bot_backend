package services_exchange_websocket

import (
	"backend/internal/enums"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_quote_repository "backend/internal/services/quote_repository/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_user "backend/internal/services/user/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
	"time"
)

type exchangeWebsocketServiceImplementation struct {
	loggerService           func() services_interface_logger.LoggerService
	symbolService           func() services_interface_symbol.SymbolService
	quoteService            func() services_interface_quote.QuoteService
	quoteRepositoryService  func() services_interface_quote_repository.QuoteRepositoryService
	exchangeService         func() services_interface_exchange.ExchangeService
	userService             func() services_interface_user.UserService
	botService              func() services_interface_bot.BotService
	dumpService             func() services_interface_dump.DumpService
	currentPriceSymbol      string
	currentPriceInterval    enums.Interval
	currentPriceStopChannel chan struct{}
	currentPriceMutex       sync.Mutex
	doneChannel             chan struct{}
	symbolsSubscriptions    map[string]chan struct{}
	symbolMutex             sync.Mutex
	reconnectDelay          time.Duration
}

func NewExchangeWebsocketService(
	loggerService func() services_interface_logger.LoggerService,
	symbolService func() services_interface_symbol.SymbolService,
	quoteService func() services_interface_quote.QuoteService,
	quoteRepositoryService func() services_interface_quote_repository.QuoteRepositoryService,
	exchangeService func() services_interface_exchange.ExchangeService,
	userService func() services_interface_user.UserService,
	botService func() services_interface_bot.BotService,
	dumpService func() services_interface_dump.DumpService,
) services_interface_exchange_websocket.ExchangeWebSocketService {
	return &exchangeWebsocketServiceImplementation{
		loggerService:           loggerService,
		symbolService:           symbolService,
		quoteService:            quoteService,
		quoteRepositoryService:  quoteRepositoryService,
		exchangeService:         exchangeService,
		userService:             userService,
		botService:              botService,
		dumpService:             dumpService,
		currentPriceStopChannel: nil,
		currentPriceMutex:       sync.Mutex{},
		doneChannel:             make(chan struct{}),
		symbolsSubscriptions:    make(map[string]chan struct{}),
		symbolMutex:             sync.Mutex{},
		reconnectDelay:          5 * time.Second,
	}
}

func (object *exchangeWebsocketServiceImplementation) Start() {
	object.loggerService().Info().Printf("starting exchange websocket service")

	go object.userData()
	go object.allMarket()
}

func (object *exchangeWebsocketServiceImplementation) Stop() {
	object.loggerService().Info().Printf("stopping exchange websocket service")

	if object.currentPriceStopChannel != nil {
		close(object.currentPriceStopChannel)
		object.currentPriceStopChannel = nil
	}

	object.symbolMutex.Lock()

	for symbol, stopChannel := range object.symbolsSubscriptions {
		object.loggerService().Info().Printf("unsubscribing from %s", symbol)
		close(stopChannel)
		delete(object.symbolsSubscriptions, symbol)
	}

	object.symbolMutex.Unlock()

	if err := object.exchangeService().DeleteListenKey(); err != nil {
		object.loggerService().Error().Printf("failed to delete listen key: %v", err)
	}

	close(object.doneChannel)
}
