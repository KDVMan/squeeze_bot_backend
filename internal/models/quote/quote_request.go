package models_quote

import (
	"backend/internal/enums"
	enums_quote "backend/internal/enums/quote"
)

type QuoteRequestModel struct {
	BotID       uint
	Symbol      string
	Interval    enums.Interval
	QuotesLimit int
	TimeEnd     int64
	Type        enums_quote.QuoteType
}
