package routes_chart_settings

import (
	routes_interface_chart_settings "backend/internal/routes/chart_settings/interface"
	services_interface_chart_settings "backend/internal/services/chart_settings/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	"github.com/go-chi/chi/v5"
)

type chartSettingsRouteImplementation struct {
	router               *chi.Mux
	loggerService        func() services_interface_logger.LoggerService
	requestService       func() services_interface_request.RequestService
	chartSettingsService func() services_interface_chart_settings.ChartSettingsService
}

func NewChartSettingsRoute(
	loggerService func() services_interface_logger.LoggerService,
	requestService func() services_interface_request.RequestService,
	chartSettingsService func() services_interface_chart_settings.ChartSettingsService,
) routes_interface_chart_settings.ChartSettingsRoute {
	route := &chartSettingsRouteImplementation{
		router:               chi.NewRouter(),
		loggerService:        loggerService,
		requestService:       requestService,
		chartSettingsService: chartSettingsService,
	}

	route.router.Get("/load", route.load())
	route.router.Post("/update", route.update())
	route.router.Get("/reset", route.reset())

	return route
}

func (object *chartSettingsRouteImplementation) GetRouter() *chi.Mux {
	return object.router
}
