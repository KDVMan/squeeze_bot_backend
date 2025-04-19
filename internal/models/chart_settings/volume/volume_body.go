package models_chart_settings_volume

type VolumeBodyData struct {
	Active bool   `json:"active"`
	Color  string `json:"color"`
}

func LoadDefaultVolumeBodyData(color string) VolumeBodyData {
	return VolumeBodyData{
		Active: true,
		Color:  color,
	}
}
