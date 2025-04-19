package models_chart_settings_ruler

type RulerCaptionBody struct {
	Up   RulerCaptionBodyData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down RulerCaptionBodyData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultRulerCaptionBody() RulerCaptionBody {
	return RulerCaptionBody{
		Up:   LoadDefaultRulerCaptionBodyData(),
		Down: LoadDefaultRulerCaptionBodyData(),
	}
}
