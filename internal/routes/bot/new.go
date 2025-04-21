package routes_bot

import (
	routes_interface_bot "backend/internal/routes/bot/interface"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type botRouteImplementation struct {
	router         *chi.Mux
	loggerService  func() services_interface_logger.LoggerService
	requestService func() services_interface_request.RequestService
	botService     func() services_interface_bot.BotService
}

func NewBotRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	botService func() services_interface_bot.BotService,
) routes_interface_bot.BotRoute {
	route := &botRouteImplementation{
		router:         chi.NewRouter(),
		loggerService:  loggerService,
		requestService: requestService,
		botService:     botService,
	}

	route.router.Post("/add", route.add())
	route.router.Post("/status", route.status())

	return route
}

func (object *botRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
