package models_chart_settings_caption

type CaptionBody struct {
	Active bool   `json:"active"`
	Color  string `json:"color"`
}

func LoadDefaultCaptionBody(active bool, color string) CaptionBody {
	return CaptionBody{
		Active: active,
		Color:  color,
	}
}
