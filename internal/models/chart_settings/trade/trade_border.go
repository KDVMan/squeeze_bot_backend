package models_chart_settings_trade

type TradeBorder struct {
	Long  TradeBorderData `gorm:"embedded;embeddedPrefix:long_" json:"long"`
	Short TradeBorderData `gorm:"embedded;embeddedPrefix:short_" json:"short"`
}

func LoadDefaultTradeBorder() TradeBorder {
	return TradeBorder{
		Long:  LoadDefaultTradeBorderData(),
		Short: LoadDefaultTradeBorderData(),
	}
}
