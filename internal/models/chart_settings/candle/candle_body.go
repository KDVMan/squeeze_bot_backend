package models_chart_settings_candle

type CandleBody struct {
	Up   CandleBodyData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down CandleBodyData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultCandleBody() CandleBody {
	return CandleBody{
		Up:   LoadDefaultCandleBodyData("#26a69aff"),
		Down: LoadDefaultCandleBodyData("#ef5350ff"),
	}
}
