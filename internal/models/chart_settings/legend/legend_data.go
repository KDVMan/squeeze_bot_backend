package models_chart_settings_legend

type LegendData struct {
	Color             string `json:"color"`
	Background        string `json:"background"`
	HorizontalPadding int    `json:"horizontalPadding"`
	VerticalPadding   int    `json:"verticalPadding"`
}

func LoadDefaultLegendData() LegendData {
	return LegendData{
		Color:             "#000000ff",
		Background:        "#ffffffff",
		HorizontalPadding: 20,
		VerticalPadding:   20,
	}
}
