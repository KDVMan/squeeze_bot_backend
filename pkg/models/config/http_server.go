package models_config

import "time"

type HttpServerModel struct {
	Address      string
	Host         string        `yaml:"host" env-required:"true"`
	Port         int           `yaml:"port" env-required:"true"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env-required:"true"`
	WriteTimeout time.Duration `yaml:"write_timeout" env-required:"true"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env-required:"true"`
}
