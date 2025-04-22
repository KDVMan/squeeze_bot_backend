package services_quote_repository

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	services_helper "backend/pkg/services/helper"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *quoteRepositoryServiceImplementation) Add(symbol string, quotesModes []*models_quote.QuoteModel) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	object.data[symbol] = quotesModes

	object.loggerService().Info().Printf("ADD, symbol: %s | len: %d\n", symbol, len(quotesModes))
}

func (object *quoteRepositoryServiceImplementation) UpdateQuote(symbol string, interval enums.Interval, trade *futures.WsAggTradeEvent) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	quotesModels, exists := object.data[symbol]
	if !exists || len(quotesModels) == 0 {
		return
	}

	price := services_helper.MustConvertStringToFloat64(trade.Price, 0, 64)
	intervalMs := enums.IntervalMilliseconds(interval)
	tradeTime := trade.TradeTime
	candleOpenTime := (tradeTime / intervalMs) * intervalMs
	candleCloseTime := candleOpenTime + intervalMs - 1

	var currentCandle *models_quote.QuoteModel

	if quotesModels[len(quotesModels)-1].TimeOpen == candleOpenTime {
		currentCandle = quotesModels[len(quotesModels)-1]
	} else {
		if len(quotesModels) >= 1440 {
			quotesModels = quotesModels[1:]
			quotesModels[len(quotesModels)-1].IsClosed = true
		}

		currentCandle = &models_quote.QuoteModel{
			Symbol:             symbol,
			Interval:           interval,
			TimeOpen:           candleOpenTime,
			TimeClose:          candleCloseTime,
			Price:              price,
			PriceOpen:          price,
			PriceHigh:          price,
			PriceLow:           price,
			PriceClose:         price,
			VolumeLeft:         0,
			Trades:             0,
			IsClosed:           false,
			TimeOpenFormatted:  services_helper.MustConvertUnixMillisecondsToString(candleOpenTime),
			TimeCloseFormatted: services_helper.MustConvertUnixMillisecondsToString(candleCloseTime),
		}

		quotesModels = append(quotesModels, currentCandle)
	}

	if price > currentCandle.PriceHigh {
		currentCandle.PriceHigh = price
	}

	if price < currentCandle.PriceLow {
		currentCandle.PriceLow = price
	}

	currentCandle.Price = price
	currentCandle.PriceClose = price

	object.data[symbol] = quotesModels

	// log.Printf(
	// 	"UPDATE, symbol: %s | time: %s | openTime: %s, closeTime: %s, closed: %v\n",
	// 	symbol,
	// 	services_helper.MustConvertUnixMillisecondsToString(trade.TradeTime),
	// 	services_helper.MustConvertUnixMillisecondsToString(candleOpenTime),
	// 	services_helper.MustConvertUnixMillisecondsToString(candleCloseTime),
	// 	currentCandle.IsClosed,
	// )
}

func (object *quoteRepositoryServiceImplementation) GetBySymbol(symbol string) []*models_quote.QuoteModel {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	quotesModels, _ := object.data[symbol]

	return quotesModels
}

// func (object *tradeRepositoryServiceImplementation) GetAll() []*models_trade.TradeModel {
// 	var tradesModels []*models_trade.TradeModel
//
// 	object.mutex.Lock()
// 	defer object.mutex.Unlock()
//
// 	if len(object.data) == 0 {
// 		return nil
// 	}
//
// 	for _, trades := range object.data {
// 		tradesModels = append(tradesModels, trades...)
// 	}
//
// 	return tradesModels
// }
