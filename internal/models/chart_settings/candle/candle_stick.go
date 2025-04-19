package models_chart_settings_candle

type CandleStick struct {
	Up   CandleStickData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down CandleStickData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultCandleStick() CandleStick {
	return CandleStick{
		Up:   LoadDefaultCandleStickData("#26a69aff"),
		Down: LoadDefaultCandleStickData("#ef5350ff"),
	}
}
