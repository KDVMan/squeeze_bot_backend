package models_chart_settings_ruler

import models_chart_settings_caption "backend/internal/models/chart_settings/caption"

type RulerInformationBody struct {
	Up   models_chart_settings_caption.CaptionBody `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down models_chart_settings_caption.CaptionBody `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultRulerInformationBody() RulerInformationBody {
	return RulerInformationBody{
		Up:   models_chart_settings_caption.LoadDefaultCaptionBody(true, "#1e88e5ff"),
		Down: models_chart_settings_caption.LoadDefaultCaptionBody(true, "#ef5350ff"),
	}
}
