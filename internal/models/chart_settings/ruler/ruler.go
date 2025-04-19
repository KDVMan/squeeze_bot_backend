package models_chart_settings_ruler

type Ruler struct {
	Body        RulerBody        `gorm:"embedded;embeddedPrefix:body_" json:"body"`
	Information RulerInformation `gorm:"embedded;embeddedPrefix:information_" json:"information"`
	Arrow       RulerArrow       `gorm:"embedded;embeddedPrefix:arrow_" json:"arrow"`
	Caption     RulerCaption     `gorm:"embedded;embeddedPrefix:caption_" json:"caption"`
}

func LoadDefaultRuler() Ruler {
	return Ruler{
		Body:        LoadDefaultRulerBody(),
		Information: LoadDefaultRulerInformation(),
		Arrow:       LoadDefaultRulerArrow(),
		Caption:     LoadDefaultRulerCaption(),
	}
}
