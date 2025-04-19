package models_chart_settings_ruler

type RulerBody struct {
	Up   RulerBodyData `gorm:"embedded;embeddedPrefix:up_" json:"up"`
	Down RulerBodyData `gorm:"embedded;embeddedPrefix:down_" json:"down"`
}

func LoadDefaultRulerBody() RulerBody {
	return RulerBody{
		Up:   LoadDefaultRulerBodyData("#1e88e534"),
		Down: LoadDefaultRulerBodyData("#ef535034"),
	}
}
