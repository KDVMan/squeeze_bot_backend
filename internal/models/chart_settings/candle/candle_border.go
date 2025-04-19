package models_chart_settings_candle

type CandleBorder struct {
	Up   CandleBorderData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down CandleBorderData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultCandleBorder() CandleBorder {
	return CandleBorder{
		Up:   LoadDefaultCandleBorderData("#26a69aff"),
		Down: LoadDefaultCandleBorderData("#ef5350ff"),
	}
}
