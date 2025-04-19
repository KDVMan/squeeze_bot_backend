package models_chart_settings_aim

import models_chart_settings_caption "backend/internal/models/chart_settings/caption"

type AimBody struct {
	Horizontal models_chart_settings_caption.CaptionBody `gorm:"embedded;embeddedPrefix:horizontal_" json:"horizontal"`
	Vertical   models_chart_settings_caption.CaptionBody `gorm:"embedded;embeddedPrefix:vertical_" json:"vertical"`
}

func LoadDefaultAimBody() AimBody {
	return AimBody{
		Horizontal: models_chart_settings_caption.LoadDefaultCaptionBody(true, "#595c6fff"),
		Vertical:   models_chart_settings_caption.LoadDefaultCaptionBody(true, "#595c6fff"),
	}
}
