package models_chart_settings_volume

type VolumeBorder struct {
	Up   VolumeBorderData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down VolumeBorderData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultVolumeBorder() VolumeBorder {
	return VolumeBorder{
		Up:   LoadDefaultVolumeBorderData("#26a69aff"),
		Down: LoadDefaultVolumeBorderData("#ef5350ff"),
	}
}
