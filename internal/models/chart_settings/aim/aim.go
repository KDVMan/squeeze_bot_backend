package models_chart_settings_aim

import enums_chart_settings "backend/internal/enums/chart_settings"

type Aim struct {
	Cursor enums_chart_settings.CursorType `json:"cursor"`
	Line   AimLine                         `gorm:"embedded;embeddedPrefix:line_" json:"line"`
	Body   AimBody                         `gorm:"embedded;embeddedPrefix:body_" json:"body"`
	Text   AimText                         `gorm:"embedded;embeddedPrefix:text_" json:"text"`
}

func LoadDefaultAim() Aim {
	return Aim{
		Cursor: enums_chart_settings.CursorTypeCrosshair,
		Line:   LoadDefaultAimLine(),
		Body:   LoadDefaultAimBody(),
		Text:   LoadDefaultAimText(),
	}
}
