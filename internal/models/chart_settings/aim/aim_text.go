package models_chart_settings_aim

import models_chart_settings_caption "backend/internal/models/chart_settings/caption"

type AimText struct {
	Horizontal models_chart_settings_caption.CaptionText `gorm:"embedded;embeddedPrefix:horizontal_" json:"horizontal"`
	Vertical   models_chart_settings_caption.CaptionText `gorm:"embedded;embeddedPrefix:vertical_" json:"vertical"`
}

func LoadDefaultAimText() AimText {
	return AimText{
		Horizontal: models_chart_settings_caption.LoadDefaultCaptionText(true, "#ffffffff"),
		Vertical:   models_chart_settings_caption.LoadDefaultCaptionText(true, "#ffffffff"),
	}
}
