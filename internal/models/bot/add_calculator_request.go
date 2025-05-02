package models_bot

import "backend/internal/enums"

type AddCalculatorRequestModel struct {
	Deposit        float64              `json:"deposit" validate:"required,gt=0"`
	Symbol         string               `json:"symbol" validate:"required"`
	Window         int64                `json:"window" validate:"gt=0"`
	Interval       enums.Interval       `json:"interval" validate:"required,interval"`
	TradeDirection enums.TradeDirection `json:"tradeDirection" validate:"required,tradeDirection"`
	Bind           enums.Bind           `json:"bind"`
	PercentIn      float64              `json:"percentIn"`
	PercentOut     float64              `json:"percentOut"`
	StopTime       int64                `json:"stopTime"`
	StopPercent    float64              `json:"stopPercent"`
	TriggerStart   float64              `json:"triggerStart" validate:"gte=0"`
	LimitQuotes    int64                `json:"limitQuotes" validate:"required,gt=0"`
}
