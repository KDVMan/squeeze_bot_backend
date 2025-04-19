package models_chart_settings_ruler

import models_chart_settings_caption "backend/internal/models/chart_settings/caption"

type RulerInformationText struct {
	Up   models_chart_settings_caption.CaptionText `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down models_chart_settings_caption.CaptionText `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultRulerInformationText() RulerInformationText {
	return RulerInformationText{
		Up:   models_chart_settings_caption.LoadDefaultCaptionText(true, "#ffffffff"),
		Down: models_chart_settings_caption.LoadDefaultCaptionText(true, "#ffffffff"),
	}
}
