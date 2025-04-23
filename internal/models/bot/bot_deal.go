package models_bot

import enums_bot "backend/internal/enums/bot"

type BotDealModel struct {
	TimeIn             int64                `json:"timeIn"`
	TimeOut            int64                `json:"timeOut"`
	PreparationPriceIn float64              `json:"preparationPriceIn"`
	PriceIn            float64              `json:"priceIn"`
	AmountIn           float64              `json:"amountIn"`
	PriceOut           float64              `json:"priceOut"`
	AmountOut          float64              `json:"amountOut"`
	IsStopTime         bool                 `json:"isStopTime"`
	IsStopPercent      bool                 `json:"isStopPercent"`
	Status             enums_bot.DealStatus `json:"status"`
}

func (object *BotDealModel) StatusIsNull() bool {
	return object.Status == enums_bot.DealStatusNull || object.Status == enums_bot.DealStatusOpenLimit || object.Status == ""
}
