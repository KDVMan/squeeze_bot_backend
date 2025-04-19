package models_chart_settings_ruler

type RulerArrow struct {
	Up   RulerArrowData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down RulerArrowData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultRulerArrow() RulerArrow {
	return RulerArrow{
		Up:   LoadDefaultRulerArrowData("#1e88e5ff"),
		Down: LoadDefaultRulerArrowData("#ef5350ff"),
	}
}
