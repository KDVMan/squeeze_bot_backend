package models_chart_settings_candle

type CandleBorderData struct {
	Active    bool   `json:"active"`
	Thickness int    `json:"thickness"`
	Color     string `json:"color"`
	MinWidth  int    `json:"minWidth"`
}

func LoadDefaultCandleBorderData(color string) CandleBorderData {
	return CandleBorderData{
		Active:    false,
		Thickness: 1,
		Color:     color,
		MinWidth:  1,
	}
}
