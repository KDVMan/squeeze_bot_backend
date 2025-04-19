package services_config

import (
	"backend/pkg/enums"
	models_config "backend/pkg/models/config"
	services_interface_config "backend/pkg/services/config/interface"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type configServiceImplementation struct {
	config *models_config.ConfigModel
}

func NewConfigService(configFile string) services_interface_config.ConfigService {
	var configModel models_config.ConfigModel

	env := os.Getenv("ENV")

	if env == "" {
		configModel.Env = enums.EnvDev
	} else {
		configModel.Env = enums.Env(env)
	}

	if err := cleanenv.ReadConfig(configFile, &configModel); err != nil {
		log.Fatalf("Failed to load config file: %s", err)
	}

	configModel.HttpServer.Address = fmt.Sprintf("%s:%d", configModel.HttpServer.Host, configModel.HttpServer.Port)

	return &configServiceImplementation{
		config: &configModel,
	}
}

func (object *configServiceImplementation) GetConfig() *models_config.ConfigModel {
	return object.config
}
