package services_chart_settings

import (
	models_chart_settings "backend/internal/models/chart_settings"
)

func (object *chartSettingsServiceImplementation) Update(request *models_chart_settings.UpdateRequestModel) (*models_chart_settings.ChartSettings, error) {
	chartSettings, err := object.Load()
	if err != nil {
		return nil, err
	}

	chartSettings.FontSize = request.FontSize
	chartSettings.BackgroundColor = request.BackgroundColor
	chartSettings.Mouse = request.Mouse
	chartSettings.Grid = request.Grid
	chartSettings.Candle = request.Candle
	chartSettings.Volume = request.Volume
	chartSettings.Legend = request.Legend
	chartSettings.Aim = request.Aim
	chartSettings.Ruler = request.Ruler
	chartSettings.Information = request.Information

	result := object.storageService().DB().Save(&chartSettings)
	if result.Error != nil {
		return nil, err
	}

	return chartSettings, nil
}
