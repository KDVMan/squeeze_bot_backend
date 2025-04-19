package models_bot

import (
	"backend/internal/enums"
)

type BotParamModel struct {
	Bind        enums.Bind `json:"bind"`
	PercentIn   float64    `json:"percentIn"`
	PercentOut  float64    `json:"percentOut"`
	StopTime    int64      `json:"stopTime"`
	StopPercent float64    `json:"stopPercent"`
}
