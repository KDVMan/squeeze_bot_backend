package models_chart_settings_ruler

type RulerCaption struct {
	Line RulerCaptionLine `gorm:"embedded;embeddedPrefix:line_" json:"line"`
	Body RulerCaptionBody `gorm:"embedded;embeddedPrefix:body_" json:"body"`
	Text RulerCaptionText `gorm:"embedded;embeddedPrefix:text_" json:"text"`
}

func LoadDefaultRulerCaption() RulerCaption {
	return RulerCaption{
		Line: LoadDefaultRulerCaptionLine(),
		Body: LoadDefaultRulerCaptionBody(),
		Text: LoadDefaultRulerCaptionText(),
	}
}
