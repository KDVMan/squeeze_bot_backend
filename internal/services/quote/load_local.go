package services_quote

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
)

func (object *quoteServiceImplementation) loadLocal(hash string, request *models_quote.QuoteRequestModel) ([]*models_quote.QuoteModel, error) {
	var quotes []*models_quote.QuoteModel
	var filteredQuotes []*models_quote.QuoteModel
	var checkTime int64
	milliseconds := enums.IntervalMilliseconds(request.Interval)

	query := object.storageService().DB().Model(&models_quote.QuoteModel{}).Where("hash = ?", hash)
	query = query.Where("time_open >= ? AND time_open <= ?", request.TimeEnd-milliseconds*int64(request.QuotesLimit), request.TimeEnd).Limit(request.QuotesLimit)

	if err := query.Order("time_open desc").Find(&quotes).Error; err != nil {
		return nil, err
	}

	for _, quote := range quotes {
		if checkTime != 0 && checkTime-quote.TimeOpen > milliseconds {
			break
		}

		filteredQuotes = append(filteredQuotes, quote)
		checkTime = quote.TimeOpen
	}

	return filteredQuotes, nil
}
