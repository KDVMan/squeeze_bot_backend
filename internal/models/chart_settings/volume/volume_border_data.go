package models_chart_settings_volume

type VolumeBorderData struct {
	Active    bool   `json:"active"`
	Thickness int    `json:"thickness"`
	Color     string `json:"color"`
}

func LoadDefaultVolumeBorderData(color string) VolumeBorderData {
	return VolumeBorderData{
		Active:    false,
		Thickness: 1,
		Color:     color,
	}
}
