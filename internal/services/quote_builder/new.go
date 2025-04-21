package services_quote_builder

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	services_interface_quote_builder "backend/internal/services/quote_builder/interface"
)

type quoteBuilderServiceImplementation struct {
	quote              *models_quote.QuoteModel
	milliseconds       int64
	millisecondsSource int64
}

func NewQuoteBuilderService(
	interval enums.Interval,
	intervalSource enums.Interval,
) services_interface_quote_builder.QuoteBuilderService {
	return &quoteBuilderServiceImplementation{
		milliseconds:       enums.IntervalMilliseconds(interval),
		millisecondsSource: enums.IntervalMilliseconds(intervalSource),
	}
}
