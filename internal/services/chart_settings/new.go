package services_chart_settings

import (
	services_interface_chart_settings "backend/internal/services/chart_settings/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type chartSettingsServiceImplementation struct {
	storageService func() services_interface_storage.StorageService
}

func NewChartSettingsService(
	storageService func() services_interface_storage.StorageService,
) services_interface_chart_settings.ChartSettingsService {
	return &chartSettingsServiceImplementation{
		storageService: storageService,
	}
}
