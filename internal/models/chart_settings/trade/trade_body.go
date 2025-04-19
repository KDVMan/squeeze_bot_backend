package models_chart_settings_trade

import "backend/internal/enums"

type TradeBody struct {
	Long  TradeBodyData `gorm:"embedded;embeddedPrefix:long_" json:"long"`
	Short TradeBodyData `gorm:"embedded;embeddedPrefix:short_" json:"short"`
}

func LoadDefaultTradeBody() TradeBody {
	return TradeBody{
		Long:  LoadDefaultTradeBodyData(enums.TradeDirectionLong),
		Short: LoadDefaultTradeBodyData(enums.TradeDirectionShort),
	}
}
