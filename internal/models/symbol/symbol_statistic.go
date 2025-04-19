package models_symbol

type SymbolStatisticModel struct {
	Price        float64 `json:"price"`
	PriceLow     float64 `json:"priceLow"`
	PriceHigh    float64 `json:"priceHigh"`
	PricePercent float64 `json:"pricePercent"`
	Volume       float64 `json:"volume"`
	Trades       int64   `json:"trades"`
}
