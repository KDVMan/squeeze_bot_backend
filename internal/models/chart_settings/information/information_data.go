package models_chart_settings_information

type InformationData struct {
	Text string          `json:"text,omitempty"`
	None InformationText `gorm:"embedded;embeddedPrefix:none_" json:"none"`
	Up   InformationText `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down InformationText `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultInformationData(text, colorNone, colorUp, colorDown string) InformationData {
	return InformationData{
		Text: text,
		None: LoadDefaultInformationText(colorNone),
		Up:   LoadDefaultInformationText(colorUp),
		Down: LoadDefaultInformationText(colorDown),
	}
}

func LoadDefaultInformationDataColor(text, color string) InformationData {
	return LoadDefaultInformationData(text, color, color, color)
}
