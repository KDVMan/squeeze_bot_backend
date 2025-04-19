package models_chart_settings_caption

import enums_chart_settings "backend/internal/enums/chart_settings"

type CaptionLine struct {
	Active    bool                          `json:"active"`
	Thickness int                           `json:"thickness"`
	Color     string                        `json:"color"`
	LineType  enums_chart_settings.LineType `json:"lineType"`
}

func LoadDefaultCaptionLine(active bool, color string, lineType enums_chart_settings.LineType) CaptionLine {
	return CaptionLine{
		Active:    active,
		Thickness: 1,
		Color:     color,
		LineType:  lineType,
	}
}
