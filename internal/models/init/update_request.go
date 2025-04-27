package models_init

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_quote "backend/internal/models/quote"
)

type UpdateRequestModel struct {
	BotID            uint                     `json:"botID" validate:"gte=0"`
	Symbol           string                   `json:"symbol" validate:"required,symbolFormat,uppercase"`
	Intervals        []*models_quote.Interval `json:"intervals,omitempty"`
	BotSortColumn    enums_bot.SortColumn     `json:"botSortColumn" validate:"required,botSortColumn"`
	BotSortDirection enums.SortDirection      `json:"botSortDirection" validate:"required,sortDirection"`
}
