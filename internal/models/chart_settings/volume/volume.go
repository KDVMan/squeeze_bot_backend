package models_chart_settings_volume

type Volume struct {
	MinLevelPercent float64      `json:"minLevelPercent"`
	MaxLevelPercent float64      `json:"maxLevelPercent"`
	Border          VolumeBorder `gorm:"embedded;embeddedPrefix:border_" json:"border"`
	Body            VolumeBody   `gorm:"embedded;embeddedPrefix:body_" json:"body"`
}

func LoadDefaultVolume() Volume {
	return Volume{
		MinLevelPercent: 15,
		MaxLevelPercent: 30,
		Border:          LoadDefaultVolumeBorder(),
		Body:            LoadDefaultVolumeBody(),
	}
}
