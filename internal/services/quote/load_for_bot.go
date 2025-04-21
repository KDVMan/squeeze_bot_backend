package services_quote

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	"time"
)

func (object *quoteServiceImplementation) LoadForBot(request *models_quote.QuoteForBotRequestModel) ([]*models_quote.QuoteModel, error) {
	var quotes []*models_quote.QuoteModel
	milliseconds := enums.IntervalMilliseconds(request.Interval)

	klines, err := object.exchangeService().Kline(request.Symbol, request.Interval.String(), request.TimeEnd, request.Limit)
	if err != nil {
		return nil, err
	}

	for _, kline := range klines {
		quote := models_quote.KlineToQuote("", request.Symbol, request.Interval, kline)

		if quote.TimeClose <= time.Now().UnixMilli() { // берем только закрытые свечи
			checkTime := (quote.TimeClose - quote.TimeOpen) + 1

			if checkTime < milliseconds {
				quote.TimeClose = quote.TimeOpen + milliseconds - 1 // битые данные (бинанс так иногда отдает)
			}
		} else {
			quote.IsClosed = false
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}
