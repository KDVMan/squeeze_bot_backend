package models_chart_settings_caption

type CaptionText struct {
	Active bool   `json:"active"`
	Color  string `json:"color"`
}

func LoadDefaultCaptionText(active bool, color string) CaptionText {
	return CaptionText{
		Active: active,
		Color:  color,
	}
}
