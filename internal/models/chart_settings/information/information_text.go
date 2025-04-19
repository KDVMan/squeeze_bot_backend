package models_chart_settings_information

type InformationText struct {
	Active bool   `json:"active"`
	Color  string `json:"color"`
}

func LoadDefaultInformationText(color string) InformationText {
	return InformationText{
		Active: true,
		Color:  color,
	}
}
