package models_chart_settings_current_price

import (
	enums_chart_settings "backend/internal/enums/chart_settings"
	models_chart_settings_caption "backend/internal/models/chart_settings/caption"
)

type CurrentPriceLine struct {
	Up   models_chart_settings_caption.CaptionLine `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down models_chart_settings_caption.CaptionLine `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultCurrentPriceLine() CurrentPriceLine {
	return CurrentPriceLine{
		Up:   models_chart_settings_caption.LoadDefaultCaptionLine(true, "#26a69aff", enums_chart_settings.LineTypeDash),
		Down: models_chart_settings_caption.LoadDefaultCaptionLine(true, "#ef5350ff", enums_chart_settings.LineTypeDash),
	}
}
