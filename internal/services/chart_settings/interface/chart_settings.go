package services_interface_chart_settings

import (
	models_chart_settings "backend/internal/models/chart_settings"
)

type ChartSettingsService interface {
	Load() (*models_chart_settings.ChartSettings, error)
	Update(*models_chart_settings.UpdateRequestModel) (*models_chart_settings.ChartSettings, error)
	Reset() (*models_chart_settings.ChartSettings, error)
}
