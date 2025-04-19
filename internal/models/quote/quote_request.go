package models_quote

import (
	"backend/internal/enums"
	enums_quote "backend/internal/enums/quote"
)

type QuoteRequestModel struct {
	Symbol      string         `json:"symbol" validate:"required,alphanum,uppercase"`
	Interval    enums.Interval `json:"interval" validate:"required,interval"`
	QuotesLimit int            `json:"quotesLimit" validate:"required,gt=0"`
	// TimeStart   int64                 `json:"-"` // нужно для загрузки calculate
	TimeEnd int64 `json:"timeEnd" validate:"gte=0"`
	// Index       int64                 `json:"index" validate:"gte=0"`
	Type enums_quote.QuoteType `json:"type" validate:"required,quoteType"`
}
