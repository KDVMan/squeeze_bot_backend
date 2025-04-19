package routes_quote

import (
	routes_interface_quote "backend/internal/routes/quote/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type quoteRouteImplementation struct {
	router         *chi.Mux
	loggerService  func() services_interface_logger.LoggerService
	requestService func() services_interface_request.RequestService
	quoteService   func() services_interface_quote.QuoteService
}

func NewQuoteRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	quoteService func() services_interface_quote.QuoteService,
) routes_interface_quote.QuoteRoute {
	route := &quoteRouteImplementation{
		router:         chi.NewRouter(),
		loggerService:  loggerService,
		requestService: requestService,
		quoteService:   quoteService,
	}

	route.router.Post("/load", route.load())

	return route
}

func (object *quoteRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
