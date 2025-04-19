package services_provider

import (
	services_chart_settings "backend/internal/services/chart_settings"
	services_interface_chart_settings "backend/internal/services/chart_settings/interface"
)

func (object *ProviderService) ChartSettingsService() services_interface_chart_settings.ChartSettingsService {
	if object.chartSettingsService == nil {
		object.chartSettingsService = services_chart_settings.NewChartSettingsService(
			object.StorageService,
		)
	}

	return object.chartSettingsService
}
