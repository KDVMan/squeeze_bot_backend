package models_chart_settings_mouse

type Mouse struct {
	WheelSpeedUp   int `json:"wheelSpeedUp" validate:"min=0"`
	WheelSpeedDown int `json:"wheelSpeedDown" validate:"min=0"`
}

func LoadDefaultMouseSettings() Mouse {
	return Mouse{
		WheelSpeedUp:   10,
		WheelSpeedDown: 10,
	}
}
