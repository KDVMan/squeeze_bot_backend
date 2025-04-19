package models_config

type BinanceModel struct {
	ApiKey            string  `yaml:"api_key" env-required:"true"`
	ApiSecret         string  `yaml:"api_secret" env-required:"true"`
	FuturesLimit      int     `yaml:"futures_limit" env-required:"true"`
	FuturesCommission float64 `yaml:"futures_commission" env-required:"true"`
}
