package services_order

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	services_helper "backend/pkg/services/helper"
)

func (object *orderServiceImplementation) getPriceIn(quote *models_quote.QuoteModel, bind enums.Bind, priceFactor float64, tickSizeFactor int) float64 {
	priceBind := object.getPriceBind(bind, quote)
	priceIn := priceBind * priceFactor

	return services_helper.Floor(priceIn, tickSizeFactor)
}

func (object *orderServiceImplementation) getPriceOut(priceIn float64, priceFactor float64, tickSizeFactor int) float64 {
	return services_helper.Floor(priceIn*priceFactor, tickSizeFactor)
}

func (object *orderServiceImplementation) getPriceStop(priceIn float64, priceFactor float64, tickSizeFactor int) float64 {
	if priceFactor <= 0 {
		return 0
	}

	return services_helper.Floor(priceIn*priceFactor, tickSizeFactor)
}

func (object *orderServiceImplementation) getPriceBind(bind enums.Bind, quote *models_quote.QuoteModel) float64 {
	switch bind {
	case enums.BindLow:
		return quote.PriceLow
	case enums.BindHigh:
		return quote.PriceHigh
	case enums.BindOpen:
		return quote.PriceOpen
	case enums.BindClose:
		return quote.PriceClose
	case enums.BindMhl:
		return (quote.PriceHigh + quote.PriceLow) / 2
	case enums.BindMoc:
		return (quote.PriceOpen + quote.PriceClose) / 2
	default:
		return 0
	}
}
