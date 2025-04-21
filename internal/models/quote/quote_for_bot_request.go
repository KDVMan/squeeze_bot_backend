package models_quote

import (
	"backend/internal/enums"
)

type QuoteForBotRequestModel struct {
	Symbol   string
	Interval enums.Interval
	Limit    int
	TimeEnd  int64
}
