package models_chart_settings_grid

type Grid struct {
	Horizontal GridData `gorm:"embedded;embeddedPrefix:horizontal_" json:"horizontal"`
	Vertical   GridData `gorm:"embedded;embeddedPrefix:vertical_" json:"vertical"`
}

func LoadDefaultGrid() Grid {
	return Grid{
		Horizontal: LoadDefaultGridData(),
		Vertical:   LoadDefaultGridData(),
	}
}
