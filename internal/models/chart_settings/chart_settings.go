package models_chart_settings

import (
	models_chart_settings_aim "backend/internal/models/chart_settings/aim"
	models_chart_settings_candle "backend/internal/models/chart_settings/candle"
	models_chart_settings_current_price "backend/internal/models/chart_settings/current_price"
	models_chart_settings_grid "backend/internal/models/chart_settings/grid"
	models_chart_settings_information "backend/internal/models/chart_settings/information"
	models_chart_settings_legend "backend/internal/models/chart_settings/legend"
	models_chart_settings_mouse "backend/internal/models/chart_settings/mouse"
	models_chart_settings_ruler "backend/internal/models/chart_settings/ruler"
	models_chart_settings_trade "backend/internal/models/chart_settings/trade"
	models_chart_settings_volume "backend/internal/models/chart_settings/volume"
)

type ChartSettings struct {
	ID              int                                              `gorm:"primaryKey" json:"-"`
	FontSize        int                                              `gorm:"not null" json:"fontSize"`
	BackgroundColor string                                           `gorm:"not null" json:"backgroundColor"`
	Mouse           models_chart_settings_mouse.Mouse                `gorm:"embedded;embeddedPrefix:mouse_" json:"mouse"`
	Grid            models_chart_settings_grid.Grid                  `gorm:"embedded;embeddedPrefix:grid_" json:"grid"`
	Candle          models_chart_settings_candle.Candle              `gorm:"embedded;embeddedPrefix:candle_" json:"candle"`
	Volume          models_chart_settings_volume.Volume              `gorm:"embedded;embeddedPrefix:volume_" json:"volume"`
	Legend          models_chart_settings_legend.Legend              `gorm:"embedded;embeddedPrefix:legend_" json:"legend"`
	CurrentPrice    models_chart_settings_current_price.CurrentPrice `gorm:"embedded;embeddedPrefix:currentPrice_" json:"currentPrice"`
	Aim             models_chart_settings_aim.Aim                    `gorm:"embedded;embeddedPrefix:aim_" json:"aim"`
	Ruler           models_chart_settings_ruler.Ruler                `gorm:"embedded;embeddedPrefix:ruler_" json:"ruler"`
	Information     models_chart_settings_information.Information    `gorm:"embedded;embeddedPrefix:information_" json:"information"`
	Trade           models_chart_settings_trade.Trade                `gorm:"embedded;embeddedPrefix:trade_" json:"trade"`
}

func (ChartSettings) TableName() string {
	return "chart_settings"
}

func LoadDefaultChartSettings() *ChartSettings {
	return &ChartSettings{
		ID:              1,
		FontSize:        14,
		BackgroundColor: "#ffffffff",
		Mouse:           models_chart_settings_mouse.LoadDefaultMouseSettings(),
		Grid:            models_chart_settings_grid.LoadDefaultGrid(),
		Candle:          models_chart_settings_candle.LoadDefaultCandle(),
		Volume:          models_chart_settings_volume.LoadDefaultVolume(),
		Legend:          models_chart_settings_legend.LoadDefaultLegend(),
		CurrentPrice:    models_chart_settings_current_price.LoadDefaultCurrentPrice(),
		Aim:             models_chart_settings_aim.LoadDefaultAim(),
		Ruler:           models_chart_settings_ruler.LoadDefaultRuler(),
		Information:     models_chart_settings_information.LoadDefaultInformation(),
		Trade:           models_chart_settings_trade.LoadDefaultTrade(),
	}
}
