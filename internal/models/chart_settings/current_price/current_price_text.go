package models_chart_settings_current_price

import models_chart_settings_caption "backend/internal/models/chart_settings/caption"

type CurrentPriceText struct {
	Up   models_chart_settings_caption.CaptionText `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down models_chart_settings_caption.CaptionText `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultCurrentPriceText() CurrentPriceText {
	return CurrentPriceText{
		Up:   models_chart_settings_caption.LoadDefaultCaptionText(true, "#ffffffff"),
		Down: models_chart_settings_caption.LoadDefaultCaptionText(true, "#ffffffff"),
	}
}
