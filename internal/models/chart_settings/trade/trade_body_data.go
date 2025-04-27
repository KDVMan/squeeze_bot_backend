package models_chart_settings_trade

import (
	"backend/internal/enums"
	enums_chart_settings "backend/internal/enums/chart_settings"
	models_chart_settings_figure "backend/internal/models/chart_settings/figure"
)

type TradeBodyData struct {
	Open  models_chart_settings_figure.FigureBody `gorm:"embedded;embeddedPrefix:open_" json:"open"`
	Close models_chart_settings_figure.FigureBody `gorm:"embedded;embeddedPrefix:close_" json:"close"`
	Stop  models_chart_settings_figure.FigureBody `gorm:"embedded;embeddedPrefix:stop_" json:"stop"`
}

func LoadDefaultTradeBodyData(tradeDirection enums.TradeDirection) TradeBodyData {
	var directionOpen, directionClose enums.Direction

	if tradeDirection == enums.TradeDirectionLong {
		directionOpen = enums.DirectionUp
		directionClose = enums.DirectionDown
	} else if tradeDirection == "short" {
		directionOpen = enums.DirectionDown
		directionClose = enums.DirectionUp
	} else {
		panic("Invalid trade type, expected 'long' or 'short'")
	}

	return TradeBodyData{
		Open:  models_chart_settings_figure.LoadDefaultFigureBody(true, enums_chart_settings.FigureTriangle, "#0cff00", 15, 20, 0, directionOpen),
		Close: models_chart_settings_figure.LoadDefaultFigureBody(true, enums_chart_settings.FigureTriangle, "#ff0000", 15, 20, 0, directionClose),
		Stop:  models_chart_settings_figure.LoadDefaultFigureBody(true, enums_chart_settings.FigureTriangle, "#0072ff", 15, 20, 0, directionClose),
	}
}
