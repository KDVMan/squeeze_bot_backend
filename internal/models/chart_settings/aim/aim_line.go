package models_chart_settings_aim

import (
	enums_chart_settings "backend/internal/enums/chart_settings"
	models_chart_settings_caption "backend/internal/models/chart_settings/caption"
)

type AimLine struct {
	Horizontal models_chart_settings_caption.CaptionLine `gorm:"embedded;embeddedPrefix:horizontal_" json:"horizontal"`
	Vertical   models_chart_settings_caption.CaptionLine `gorm:"embedded;embeddedPrefix:vertical_" json:"vertical"`
}

func LoadDefaultAimLine() AimLine {
	return AimLine{
		Horizontal: models_chart_settings_caption.LoadDefaultCaptionLine(true, "#808292ff", enums_chart_settings.LineTypeDash),
		Vertical:   models_chart_settings_caption.LoadDefaultCaptionLine(true, "#808292ff", enums_chart_settings.LineTypeDash),
	}
}
