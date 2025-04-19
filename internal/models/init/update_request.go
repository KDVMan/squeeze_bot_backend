package models_init

import (
	models_quote "backend/internal/models/quote"
)

type UpdateRequestModel struct {
	Symbol    string                   `json:"symbol" validate:"required,symbolFormat,uppercase"`
	Intervals []*models_quote.Interval `json:"intervals,omitempty"`
}
