package models_config

type LoggerModel struct {
	UseFileOnDev  bool   `yaml:"use_file_on_dev"`
	UseFileOnProd bool   `yaml:"use_file_on_prod"`
	ErrorFileName string `yaml:"error_file_name" env-required:"true"`
	InfoFileName  string `yaml:"info_file_name" env-required:"true"`
}
