package routes_symbol_list

import (
	routes_interface_symbol_list "backend/internal/routes/symbol_list/interface"
	services_interface_symbol_list "backend/internal/services/symbol_list/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type symbolListRouteImplementation struct {
	router            *chi.Mux
	loggerService     func() services_interface_logger.LoggerService
	requestService    func() services_interface_request.RequestService
	symbolListService func() services_interface_symbol_list.SymbolListService
}

func NewSymbolListRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	symbolListService func() services_interface_symbol_list.SymbolListService,
) routes_interface_symbol_list.SymbolListRoute {
	route := &symbolListRouteImplementation{
		router:            chi.NewRouter(),
		loggerService:     loggerService,
		requestService:    requestService,
		symbolListService: symbolListService,
	}

	route.router.Get("/load", route.load())
	route.router.Post("/update", route.update())

	return route
}

func (object *symbolListRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
