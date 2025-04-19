package models_chart_settings_trade

type Trade struct {
	Border TradeBorder `gorm:"embedded;embeddedPrefix:border_" json:"border"`
	Body   TradeBody   `gorm:"embedded;embeddedPrefix:body_" json:"body"`
}

func LoadDefaultTrade() Trade {
	return Trade{
		Border: LoadDefaultTradeBorder(),
		Body:   LoadDefaultTradeBody(),
	}
}
