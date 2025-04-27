package models_quote

import (
	models_bot "backend/internal/models/bot"
	models_deal "backend/internal/models/deal"
)

type QuoteResponseModel struct {
	Quotes   []*QuoteModel            `json:"quotes"`
	TimeFrom int64                    `json:"timeFrom"`
	TimeTo   int64                    `json:"timeTo"`
	Bot      *models_bot.BotModel     `json:"bot"`
	Deals    []*models_deal.DealModel `json:"deals"`
}
