package models_symbol

type SymbolLimitModel struct {
	LeftMin       float64 `json:"leftMin"`
	LeftMax       float64 `json:"leftMax"`
	LeftStep      float64 `json:"leftStep"`
	LeftPrecision int     `json:"leftPrecision"`
	RightMin      float64 `json:"rightMin"`
	RightMax      float64 `json:"rightMax"`
	Precision     int     `json:"precision"`
	TickSize      float64 `json:"tickSize"`
}
