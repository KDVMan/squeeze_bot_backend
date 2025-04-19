package models_chart_settings_figure

type FigureBorder struct {
	Active    bool   `json:"active"`
	Color     string `json:"color"`
	Thickness int    `json:"thickness"`
}

func LoadDefaultFigureBorder(active bool, color string, thickness int) FigureBorder {
	return FigureBorder{
		Active:    active,
		Color:     color,
		Thickness: thickness,
	}
}
