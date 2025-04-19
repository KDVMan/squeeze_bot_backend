package models_config

import "backend/pkg/enums"

type ConfigModel struct {
	Env        enums.Env
	Logger     LoggerModel     `yaml:"logger"`
	DB         DBModel         `yaml:"db"`
	HttpServer HttpServerModel `yaml:"http_server"`
	Binance    BinanceModel    `yaml:"binance"`
}
