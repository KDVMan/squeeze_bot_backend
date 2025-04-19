package models_chart_settings_ruler

import models_chart_settings_caption "backend/internal/models/chart_settings/caption"

type RulerCaptionTextData struct {
	Right  models_chart_settings_caption.CaptionText `gorm:"embedded;embeddedPrefix:right_" json:"right"`
	Bottom models_chart_settings_caption.CaptionText `gorm:"embedded;embeddedPrefix:bottom_" json:"bottom"`
}

func LoadDefaultRulerCaptionTextData() RulerCaptionTextData {
	return RulerCaptionTextData{
		Right:  models_chart_settings_caption.LoadDefaultCaptionText(false, "#ffffffff"),
		Bottom: models_chart_settings_caption.LoadDefaultCaptionText(false, "#ffffffff"),
	}
}
