package models_bot

import "backend/internal/enums"

type AddRequestModel struct {
	Deposit        float64              `json:"deposit" validate:"required,gt=0"`
	IsReal         bool                 `json:"isReal"`
	Symbol         string               `json:"symbol" validate:"required"`
	Window         int64                `json:"window" validate:"gte=0"`
	Interval       enums.Interval       `json:"interval" validate:"required,interval"`
	TradeDirection enums.TradeDirection `json:"tradeDirection" validate:"required,tradeDirection"`
	Bind           enums.Bind           `json:"bind" validate:"required,bind"`
	PercentIn      float64              `json:"percentIn" validate:"required,gt=0"`
	PercentOut     float64              `json:"percentOut" validate:"required,gt=0"`
	StopTime       int64                `json:"stopTime" validate:"gte=0"`
	StopPercent    float64              `json:"stopPercent" validate:"gte=0"`
	TriggerStart   float64              `json:"triggerStart" validate:"gte=0,ltefield=PercentIn"`
	LimitQuotes    int64                `json:"limitQuotes" validate:"required,gt=0"`
	IsCalculator   bool                 `json:"isCalculator"`
}
