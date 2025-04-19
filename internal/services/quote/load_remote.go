package services_quote

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	"gorm.io/gorm/clause"
	"time"
)

func (object *quoteServiceImplementation) loadRemote(hash string, request *models_quote.QuoteRequestModel) ([]*models_quote.QuoteModel, error) {
	var quotes []*models_quote.QuoteModel
	milliseconds := enums.IntervalMilliseconds(request.Interval)

	klines, err := object.exchangeService().Kline(request.Symbol, string(request.Interval), request.TimeEnd, request.QuotesLimit)
	if err != nil {
		return nil, err
	}

	tx := object.storageService().DB().Begin()
	if tx.Error != nil {
		return nil, err
	}

	for _, kline := range klines {
		quote := models_quote.KlineToQuote(hash, request.Symbol, request.Interval, kline)

		if quote.TimeClose <= time.Now().UnixMilli() { // пишем в базу только закрытые свечи
			checkTime := (quote.TimeClose - quote.TimeOpen) + 1

			if checkTime < milliseconds {
				quote.TimeClose = quote.TimeOpen + milliseconds - 1 // битые данные (бинанс так иногда отдает)
			}

			if err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&quote).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}

		quotes = append(quotes, quote)
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return quotes, nil
}
