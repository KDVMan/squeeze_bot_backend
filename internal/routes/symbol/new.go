package routes_symbol

import (
	routes_interface_symbol "backend/internal/routes/symbol/interface"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type symbolRouteImplementation struct {
	router          *chi.Mux
	loggerService   func() services_interface_logger.LoggerService
	requestService  func() services_interface_request.RequestService
	symbolService   func() services_interface_symbol.SymbolService
	exchangeService func() services_interface_exchange.ExchangeService
}

func NewSymbolRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	symbolService func() services_interface_symbol.SymbolService,
	exchangeService func() services_interface_exchange.ExchangeService,
) routes_interface_symbol.SymbolRoute {
	route := &symbolRouteImplementation{
		router:          chi.NewRouter(),
		loggerService:   loggerService,
		requestService:  requestService,
		symbolService:   symbolService,
		exchangeService: exchangeService,
	}

	route.router.Get("/download", route.download())
	route.router.Post("/search", route.search())

	return route
}

func (object *symbolRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
