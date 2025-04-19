package models_quote

import (
	"backend/internal/enums"
	services_helper "backend/pkg/services/helper"
)

type QuotePercentModel struct {
	Low  float64 `json:"low"`
	Body float64 `json:"body"`
	High float64 `json:"high"`
}

func GetPercent(direction enums.Direction, priceOpen float64, priceHigh float64, priceLow float64, priceClose float64, fix int) QuotePercentModel {
	if direction == enums.DirectionUp {
		return QuotePercentModel{
			Low:  services_helper.GetPercentFromMinMax(priceLow, priceOpen, fix),
			Body: services_helper.GetPercentFromMinMax(priceOpen, priceClose, fix),
			High: services_helper.GetPercentFromMinMax(priceClose, priceHigh, fix),
		}
	} else {
		return QuotePercentModel{
			High: services_helper.GetPercentFromMinMax(priceHigh, priceOpen, fix),
			Body: services_helper.GetPercentFromMinMax(priceOpen, priceClose, fix),
			Low:  services_helper.GetPercentFromMinMax(priceClose, priceLow, fix),
		}
	}
}
