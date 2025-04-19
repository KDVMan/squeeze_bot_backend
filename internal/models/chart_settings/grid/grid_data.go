package models_chart_settings_grid

import enums_chart_settings "backend/internal/enums/chart_settings"

type GridData struct {
	Active    bool                          `json:"active"`
	Thickness int                           `json:"thickness"`
	Color     string                        `json:"color"`
	LineType  enums_chart_settings.LineType `json:"lineType"`
}

func LoadDefaultGridData() GridData {
	return GridData{
		Active:    true,
		Thickness: 1,
		Color:     "#efefefff",
		LineType:  enums_chart_settings.LineTypeSolid,
	}
}
