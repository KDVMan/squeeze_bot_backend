package services_provider

import (
	routes_chart_settings "backend/internal/routes/chart_settings"
	routes_interface_chart_settings "backend/internal/routes/chart_settings/interface"
)

func (object *ProviderService) ChartSettingsRoute() routes_interface_chart_settings.ChartSettingsRoute {
	if object.chartSettingsRoute == nil {
		object.chartSettingsRoute = routes_chart_settings.NewChartSettingsRoute(
			object.LoggerService,
			object.RequestService,
			object.ChartSettingsService,
		)
	}

	return object.chartSettingsRoute
}
