package models_chart_settings_ruler

import (
	enums_chart_settings "backend/internal/enums/chart_settings"
	models_chart_settings_caption "backend/internal/models/chart_settings/caption"
)

type RulerCaptionLineData struct {
	Horizontal models_chart_settings_caption.CaptionLine `gorm:"embedded;embeddedPrefix:horizontal_" json:"horizontal"`
	Vertical   models_chart_settings_caption.CaptionLine `gorm:"embedded;embeddedPrefix:vertical_" json:"vertical"`
}

func LoadDefaultRulerCaptionLineData() RulerCaptionLineData {
	return RulerCaptionLineData{
		Horizontal: models_chart_settings_caption.LoadDefaultCaptionLine(false, "#000000ff", enums_chart_settings.LineTypeSolid),
		Vertical:   models_chart_settings_caption.LoadDefaultCaptionLine(false, "#000000ff", enums_chart_settings.LineTypeSolid),
	}
}
