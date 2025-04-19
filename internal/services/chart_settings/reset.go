package services_chart_settings

import (
	models_chart_settings "backend/internal/models/chart_settings"
)

func (object *chartSettingsServiceImplementation) Reset() (*models_chart_settings.ChartSettings, error) {
	var chartSettings = models_chart_settings.LoadDefaultChartSettings()

	if err := object.storageService().DB().Save(chartSettings).Error; err != nil {
		return nil, err
	}

	return chartSettings, nil
}
