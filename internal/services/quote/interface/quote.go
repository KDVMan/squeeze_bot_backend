package services_interface_quote

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	"github.com/adshao/go-binance/v2/futures"
)

type QuoteService interface {
	Load(*models_quote.QuoteRequestModel) (*models_quote.QuoteResponseModel, error)
	LoadRange(string, *models_quote.QuoteRangeModel) ([]*models_quote.QuoteModel, error)
	LoadForBot(*models_quote.QuoteForBotRequestModel) ([]*models_quote.QuoteModel, error)
	CurrentPrice(string, enums.Interval, futures.WsKline) error
}
