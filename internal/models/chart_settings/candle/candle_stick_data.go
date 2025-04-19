package models_chart_settings_candle

type CandleStickData struct {
	Active    bool   `json:"active"`
	Thickness int    `json:"thickness"`
	Color     string `json:"color"`
}

func LoadDefaultCandleStickData(color string) CandleStickData {
	return CandleStickData{
		Active:    true,
		Thickness: 1,
		Color:     color,
	}
}
