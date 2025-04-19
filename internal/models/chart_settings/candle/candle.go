package models_chart_settings_candle

type Candle struct {
	MinOnScreen int          `json:"minOnScreen"`
	MaxOnScreen int          `json:"maxOnScreen"`
	GapSize     int          `json:"gapSize"`
	Stick       CandleStick  `gorm:"embedded;embeddedPrefix:stick_" json:"stick"`
	Border      CandleBorder `gorm:"embedded;embeddedPrefix:border_" json:"border"`
	Body        CandleBody   `gorm:"embedded;embeddedPrefix:body_" json:"body"`
}

func LoadDefaultCandle() Candle {
	return Candle{
		MinOnScreen: 5,
		MaxOnScreen: 2000,
		GapSize:     1,
		Stick:       LoadDefaultCandleStick(),
		Border:      LoadDefaultCandleBorder(),
		Body:        LoadDefaultCandleBody(),
	}
}
