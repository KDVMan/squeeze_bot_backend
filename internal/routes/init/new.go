package routes_init

import (
	routes_interface_init "backend/internal/routes/init/interface"
	services_interface_init "backend/internal/services/init/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type initRouteImplementation struct {
	router         *chi.Mux
	loggerService  func() services_interface_logger.LoggerService
	requestService func() services_interface_request.RequestService
	initService    func() services_interface_init.InitService
}

func NewInitRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	initService func() services_interface_init.InitService,
) routes_interface_init.InitRoute {
	route := &initRouteImplementation{
		router:         chi.NewRouter(),
		loggerService:  loggerService,
		requestService: requestService,
		initService:    initService,
	}

	route.router.Get("/load", route.load())
	route.router.Get("/load_variable", route.loadVariable())
	route.router.Post("/update", route.update())

	return route
}

func (object *initRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
