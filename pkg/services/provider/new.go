package services_provider

import (
	routes_interface_bot "backend/internal/routes/bot/interface"
	routes_interface_chart_settings "backend/internal/routes/chart_settings/interface"
	routes_interface_init "backend/internal/routes/init/interface"
	routes_interface_quote "backend/internal/routes/quote/interface"
	routes_interface_symbol "backend/internal/routes/symbol/interface"
	routes_interface_symbol_list "backend/internal/routes/symbol_list/interface"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_bot_repository "backend/internal/services/bot_repository/interface"
	services_interface_chart_settings "backend/internal/services/chart_settings/interface"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
	services_interface_exchange_order "backend/internal/services/exchange_order/interface"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
	services_interface_init "backend/internal/services/init/interface"
	services_interface_order "backend/internal/services/order/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_quote_repository "backend/internal/services/quote_repository/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_symbol_list "backend/internal/services/symbol_list/interface"
	services_interface_user "backend/internal/services/user/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_exchange_binance "backend/pkg/services/exchange_binance/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	services_interface_router "backend/pkg/services/router/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
	"context"
)

type ProviderService struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	// ядро
	configService          services_interface_config.ConfigService
	loggerService          services_interface_logger.LoggerService
	storageService         services_interface_storage.StorageService
	dumpService            services_interface_dump.DumpService
	routerService          services_interface_router.RouterService
	requestService         services_interface_request.RequestService
	exchangeBinanceService services_interface_exchange_binance.ExchangeBinanceService

	// роуты
	initRoute          routes_interface_init.InitRoute
	symbolRoute        routes_interface_symbol.SymbolRoute
	symbolListRoute    routes_interface_symbol_list.SymbolListRoute
	quoteRoute         routes_interface_quote.QuoteRoute
	chartSettingsRoute routes_interface_chart_settings.ChartSettingsRoute
	botRoute           routes_interface_bot.BotRoute

	// сервисы
	websocketService         services_interface_websocket.WebsocketService
	initService              services_interface_init.InitService
	symbolService            services_interface_symbol.SymbolService
	symbolListService        services_interface_symbol_list.SymbolListService
	exchangeService          services_interface_exchange.ExchangeService
	exchangeLimitService     services_interface_exchange_limit.ExchangeLimitService
	exchangeWebsocketService services_interface_exchange_websocket.ExchangeWebSocketService
	exchangeOrderService     services_interface_exchange_order.ExchangeOrderService
	quoteService             services_interface_quote.QuoteService
	quoteRepositoryService   services_interface_quote_repository.QuoteRepositoryService
	chartSettingsService     services_interface_chart_settings.ChartSettingsService
	userService              services_interface_user.UserService
	botService               services_interface_bot.BotService
	botRepositoryService     services_interface_bot_repository.BotRepositoryService
	orderService             services_interface_order.OrderService
}

func NewProviderService(parentCtx context.Context) *ProviderService {
	ctx, cancelCtx := context.WithCancel(parentCtx)

	return &ProviderService{
		ctx:       ctx,
		cancelCtx: cancelCtx,
	}
}

func (object *ProviderService) Shutdown() {
	object.loggerService.Info().Println("shutting down provider service...")

	if object.exchangeOrderService != nil {
		object.exchangeOrderService.Stop()
	}

	if object.exchangeWebsocketService != nil {
		object.exchangeWebsocketService.Stop()
	}

	if object.websocketService != nil {
		object.websocketService.Stop()
	}

	object.cancelCtx()

	object.loggerService.Info().Println("provider service stopped.")
}
