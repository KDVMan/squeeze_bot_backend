package models_config

type DBModel struct {
	Path string `yaml:"path" env-required:"true"`
}
