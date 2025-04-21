package services_interface_quote_repository

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	"github.com/adshao/go-binance/v2/futures"
)

type QuoteRepositoryService interface {
	Add(string, []*models_quote.QuoteModel)
	UpdateQuote(string, enums.Interval, *futures.WsAggTradeEvent)
	GetBySymbol(string) []*models_quote.QuoteModel
}
