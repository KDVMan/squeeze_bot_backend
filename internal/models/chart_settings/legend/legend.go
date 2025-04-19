package models_chart_settings_legend

type Legend struct {
	Horizontal LegendData `gorm:"embedded;embeddedPrefix:horizontal_" json:"horizontal"`
	Vertical   LegendData `gorm:"embedded;embeddedPrefix:vertical_" json:"vertical"`
}

func LoadDefaultLegend() Legend {
	return Legend{
		Horizontal: LoadDefaultLegendData(),
		Vertical:   LoadDefaultLegendData(),
	}
}
