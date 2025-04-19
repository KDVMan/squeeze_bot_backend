package models_chart_settings_volume

type VolumeBody struct {
	Up   VolumeBodyData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down VolumeBodyData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultVolumeBody() VolumeBody {
	return VolumeBody{
		Up:   LoadDefaultVolumeBodyData("#3a785fff"),
		Down: LoadDefaultVolumeBodyData("#953745ff"),
	}
}
