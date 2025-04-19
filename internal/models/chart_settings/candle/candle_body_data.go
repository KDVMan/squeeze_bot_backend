package models_chart_settings_candle

type CandleBodyData struct {
	Active   bool   `json:"active"`
	Color    string `json:"color"`
	MinWidth int    `json:"minWidth"`
}

func LoadDefaultCandleBodyData(color string) CandleBodyData {
	return CandleBodyData{
		Active:   true,
		Color:    color,
		MinWidth: 1,
	}
}
