package models_chart_settings_figure

import (
	"backend/internal/enums"
	enums_chart_settings "backend/internal/enums/chart_settings"
)

type FigureBody struct {
	Active    bool                        `json:"active"`
	Figure    enums_chart_settings.Figure `json:"figure"`
	Color     string                      `json:"color"`
	Width     int                         `json:"width"`
	Height    int                         `json:"height"`
	Thickness int                         `json:"thickness"`
	Direction enums.Direction             `json:"direction"`
}

func LoadDefaultFigureBody(active bool, figure enums_chart_settings.Figure, color string, width, height, thickness int, direction enums.Direction) FigureBody {
	return FigureBody{
		Active:    active,
		Figure:    figure,
		Color:     color,
		Width:     width,
		Height:    height,
		Thickness: thickness,
		Direction: direction,
	}
}
