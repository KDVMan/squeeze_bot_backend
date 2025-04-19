package models_chart_settings

import (
	models_chart_settings_aim "backend/internal/models/chart_settings/aim"
	models_chart_settings_candle "backend/internal/models/chart_settings/candle"
	models_chart_settings_grid "backend/internal/models/chart_settings/grid"
	models_chart_settings_information "backend/internal/models/chart_settings/information"
	models_chart_settings_legend "backend/internal/models/chart_settings/legend"
	models_chart_settings_mouse "backend/internal/models/chart_settings/mouse"
	models_chart_settings_ruler "backend/internal/models/chart_settings/ruler"
	models_chart_settings_volume "backend/internal/models/chart_settings/volume"
)

type UpdateRequestModel struct {
	FontSize        int                                           `validate:"required,gt=0"`
	BackgroundColor string                                        `validate:"required,hexcolor"`
	Mouse           models_chart_settings_mouse.Mouse             `validate:"required"`
	Grid            models_chart_settings_grid.Grid               `validate:"required"`
	Candle          models_chart_settings_candle.Candle           `validate:"required"`
	Volume          models_chart_settings_volume.Volume           `validate:"required"`
	Legend          models_chart_settings_legend.Legend           `validate:"required"`
	Aim             models_chart_settings_aim.Aim                 `validate:"required"`
	Ruler           models_chart_settings_ruler.Ruler             `validate:"required"`
	Information     models_chart_settings_information.Information `validate:"required"`
}
