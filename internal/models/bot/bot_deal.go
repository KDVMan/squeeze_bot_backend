package models_bot

import enums_bot "backend/internal/enums/bot"

type BotDealModel struct {
	TimeIn             int64                `json:"timeIn"`
	TimeOut            int64                `json:"timeOut"`
	PriceIn            float64              `json:"priceIn"`
	PriceOut           float64              `json:"priceOut"`
	IsStopTime         bool                 `json:"isStopTime"`
	IsStopPercent      bool                 `json:"isStopPercent"`
	CalculateTimeOut   int64                `json:"calculateTimeOut"`
	CalculatePriceOut  float64              `json:"calculatePriceOut"`
	CalculatePriceStop float64              `json:"calculatePriceStop"`
	Status             enums_bot.DealStatus `json:"status"`
}

func (object *BotDealModel) IsNull() bool {
	return object.Status == enums_bot.DealStatusNull || object.Status == ""
}
