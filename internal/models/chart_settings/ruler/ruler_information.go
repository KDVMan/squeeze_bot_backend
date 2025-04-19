package models_chart_settings_ruler

type RulerInformation struct {
	Body              RulerInformationBody `gorm:"embedded;embeddedPrefix:body_" json:"body"`
	Text              RulerInformationText `gorm:"embedded;embeddedPrefix:text_" json:"text"`
	HorizontalPadding int                  `json:"horizontalPadding"`
	VerticalPadding   int                  `json:"verticalPadding"`
	Indent            int                  `json:"indent"`
}

func LoadDefaultRulerInformation() RulerInformation {
	return RulerInformation{
		Body:              LoadDefaultRulerInformationBody(),
		Text:              LoadDefaultRulerInformationText(),
		HorizontalPadding: 10,
		VerticalPadding:   10,
		Indent:            10,
	}
}
