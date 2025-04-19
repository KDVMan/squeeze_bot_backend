package models_chart_settings_current_price

import models_chart_settings_caption "backend/internal/models/chart_settings/caption"

type CurrentPriceBody struct {
	Up   models_chart_settings_caption.CaptionBody `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down models_chart_settings_caption.CaptionBody `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultCurrentPriceBody() CurrentPriceBody {
	return CurrentPriceBody{
		Up:   models_chart_settings_caption.LoadDefaultCaptionBody(true, "#26a69aff"),
		Down: models_chart_settings_caption.LoadDefaultCaptionBody(true, "#ef5350ff"),
	}
}
