package models_chart_settings_current_price

type CurrentPrice struct {
	Line CurrentPriceLine `gorm:"embedded;embeddedPrefix:line_" json:"line"`
	Body CurrentPriceBody `gorm:"embedded;embeddedPrefix:body_" json:"body"`
	Text CurrentPriceText `gorm:"embedded;embeddedPrefix:text_" json:"text"`
}

func LoadDefaultCurrentPrice() CurrentPrice {
	return CurrentPrice{
		Line: LoadDefaultCurrentPriceLine(),
		Body: LoadDefaultCurrentPriceBody(),
		Text: LoadDefaultCurrentPriceText(),
	}
}
