package models_chart_settings_ruler

type RulerCaptionText struct {
	Up   RulerCaptionTextData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down RulerCaptionTextData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultRulerCaptionText() RulerCaptionText {
	return RulerCaptionText{
		Up:   LoadDefaultRulerCaptionTextData(),
		Down: LoadDefaultRulerCaptionTextData(),
	}
}
