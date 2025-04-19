package models_chart_settings_trade

import models_chart_settings_figure "backend/internal/models/chart_settings/figure"

type TradeBorderData struct {
	Open  models_chart_settings_figure.FigureBorder `gorm:"embedded;embeddedPrefix:open_" json:"open"`
	Close models_chart_settings_figure.FigureBorder `gorm:"embedded;embeddedPrefix:close_" json:"close"`
	Stop  models_chart_settings_figure.FigureBorder `gorm:"embedded;embeddedPrefix:stop_" json:"stop"`
}

func LoadDefaultTradeBorderData() TradeBorderData {
	return TradeBorderData{
		Open:  models_chart_settings_figure.LoadDefaultFigureBorder(true, "#000000ff", 1),
		Close: models_chart_settings_figure.LoadDefaultFigureBorder(true, "#000000ff", 1),
		Stop:  models_chart_settings_figure.LoadDefaultFigureBorder(true, "#000000ff", 1),
	}
}
