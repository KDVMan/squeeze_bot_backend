package services_router

import (
	routes_interface_chart_settings "backend/internal/routes/chart_settings/interface"
	routes_interface_init "backend/internal/routes/init/interface"
	routes_interface_quote "backend/internal/routes/quote/interface"
	routes_interface_symbol "backend/internal/routes/symbol/interface"
	routes_interface_symbol_list "backend/internal/routes/symbol_list/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_router "backend/pkg/services/router/interface"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

type routerServiceImplementation struct {
	router           *chi.Mux
	loggerService    func() services_interface_logger.LoggerService
	websocketService func() services_interface_websocket.WebsocketService
}

func NewRouterService(
	loggerService func() services_interface_logger.LoggerService,
	websocketService func() services_interface_websocket.WebsocketService,
	initRoute func() routes_interface_init.InitRoute,
	symbolRoute func() routes_interface_symbol.SymbolRoute,
	symbolListRoute func() routes_interface_symbol_list.SymbolListRoute,
	quoteRoute func() routes_interface_quote.QuoteRoute,
	chartSettingsRoute func() routes_interface_chart_settings.ChartSettingsRoute,
	// tradeRoute func() routes_trade_interface.TradeRoute,
) services_interface_router.RouterService {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	routerService := &routerServiceImplementation{
		router:           router,
		loggerService:    loggerService,
		websocketService: websocketService,
	}

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		routerService.websocket(w, r)
	})

	router.Route("/api", func(r chi.Router) {
		r.Mount("/init", initRoute().GetRouter())
		r.Mount("/symbol", symbolRoute().GetRouter())
		r.Mount("/symbol_list", symbolListRoute().GetRouter())
		r.Mount("/quote", quoteRoute().GetRouter())
		r.Mount("/chart_settings", chartSettingsRoute().GetRouter())
		// r.Mount("/trade", tradeRoute().GetRouter())
	})

	return routerService
}

func (object *routerServiceImplementation) GetRouter() *chi.Mux {
	return object.router
}
