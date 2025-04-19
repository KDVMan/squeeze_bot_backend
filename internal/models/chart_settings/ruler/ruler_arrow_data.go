package models_chart_settings_ruler

type RulerArrowData struct {
	Active bool   `json:"active"`
	Color  string `json:"color"`
}

func LoadDefaultRulerArrowData(color string) RulerArrowData {
	return RulerArrowData{
		Active: true,
		Color:  color,
	}
}
