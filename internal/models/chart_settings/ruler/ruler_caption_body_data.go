package models_chart_settings_ruler

import models_chart_settings_caption "backend/internal/models/chart_settings/caption"

type RulerCaptionBodyData struct {
	Right  models_chart_settings_caption.CaptionBody `gorm:"embedded;embeddedPrefix:right_" json:"right"`
	Bottom models_chart_settings_caption.CaptionBody `gorm:"embedded;embeddedPrefix:bottom_" json:"bottom"`
}

func LoadDefaultRulerCaptionBodyData() RulerCaptionBodyData {
	return RulerCaptionBodyData{
		Right:  models_chart_settings_caption.LoadDefaultCaptionBody(false, "#000000ff"),
		Bottom: models_chart_settings_caption.LoadDefaultCaptionBody(false, "#000000ff"),
	}
}
