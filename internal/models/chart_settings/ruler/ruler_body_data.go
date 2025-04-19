package models_chart_settings_ruler

type RulerBodyData struct {
	Active bool   `json:"active"`
	Color  string `json:"color"`
}

func LoadDefaultRulerBodyData(color string) RulerBodyData {
	return RulerBodyData{
		Active: true,
		Color:  color,
	}
}
