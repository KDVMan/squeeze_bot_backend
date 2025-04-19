package services_chart_settings

import (
	models_chart_settings "backend/internal/models/chart_settings"
	"errors"
	"gorm.io/gorm"
)

func (object *chartSettingsServiceImplementation) Load() (*models_chart_settings.ChartSettings, error) {
	var chartSettings *models_chart_settings.ChartSettings

	if err := object.storageService().DB().First(&chartSettings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			chartSettings = models_chart_settings.LoadDefaultChartSettings()

			if err = object.storageService().DB().Create(chartSettings).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return chartSettings, nil
}
