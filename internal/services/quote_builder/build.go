package services_quote_builder

import (
	models_quote "backend/internal/models/quote"
)

func (object *quoteBuilderServiceImplementation) Build(quote *models_quote.QuoteModel) *models_quote.QuoteModel {
	if object.milliseconds == object.millisecondsSource {
		return quote
	}

	currentInterval := quote.TimeOpen / object.milliseconds

	if object.quote != nil && currentInterval != (object.quote.TimeOpen/object.milliseconds) {
		if object.isLast(quote, currentInterval) {
			object.quote = nil
		} else {
			newQuote := *quote
			newQuote.IsClosed = false
			object.quote = &newQuote
		}

		return nil
	}

	if object.quote == nil {
		newQuote := *quote
		newQuote.IsClosed = false
		object.quote = &newQuote
	} else {
		object.quote.TimeClose = quote.TimeClose
		object.quote.PriceClose = quote.PriceClose
		object.quote.VolumeLeft += quote.VolumeLeft
		object.quote.VolumeRight += quote.VolumeRight
		// object.quote.Trades += quote.Trades
		// object.quote.VolumeBuyLeft += quote.VolumeBuyLeft
		// object.quote.VolumeBuyRight += quote.VolumeBuyRight

		if quote.PriceOpen > 0 {
			object.quote.PriceHigh = max(quote.PriceHigh, object.quote.PriceHigh)
			object.quote.PriceLow = min(quote.PriceLow, object.quote.PriceLow)
		} else {
			object.quote.PriceHigh = min(quote.PriceHigh, object.quote.PriceHigh)
			object.quote.PriceLow = max(quote.PriceLow, object.quote.PriceLow)
		}
	}

	if object.isLast(quote, currentInterval) {
		newQuote := *object.quote
		newQuote.IsClosed = true
		object.quote = nil

		return &newQuote
	}

	return object.quote
}

func (object *quoteBuilderServiceImplementation) isLast(quote *models_quote.QuoteModel, currentInterval int64) bool {
	return ((quote.TimeOpen + object.millisecondsSource) / object.milliseconds) != currentInterval
}
