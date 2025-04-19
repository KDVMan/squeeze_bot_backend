package services_interface_config

import models_config "backend/pkg/models/config"

type ConfigService interface {
	GetConfig() *models_config.ConfigModel
}
