package models_chart_settings_ruler

type RulerCaptionLine struct {
	Up   RulerCaptionLineData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down RulerCaptionLineData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultRulerCaptionLine() RulerCaptionLine {
	return RulerCaptionLine{
		Up:   LoadDefaultRulerCaptionLineData(),
		Down: LoadDefaultRulerCaptionLineData(),
	}
}
