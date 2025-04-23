package models_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	"backend/internal/models"
)

type BotModel struct {
	models.DbModelWithID
	Hash            string               `gorm:"uniqueIndex:unique_bot_01;not null" json:"-"`
	Deposit         float64              `json:"deposit"`
	IsReal          bool                 `json:"isReal"`
	Symbol          string               `json:"symbol"`
	Interval        enums.Interval       `json:"interval"`
	TradeDirection  enums.TradeDirection `json:"tradeDirection"`
	Window          int                  `gorm:"default:0" json:"window"`
	LimitQuotes     int64                `json:"limitQuotes"`
	PrevParam       BotParamModel        `gorm:"embedded;embeddedPrefix:prev_param_" json:"prevParam"`
	CurrentParam    BotParamModel        `gorm:"embedded;embeddedPrefix:current_param_" json:"currentParam"`
	NextParam       BotParamModel        `gorm:"embedded;embeddedPrefix:next_param_" json:"nextParam"`
	Multiplier      BotMultiplierModel   `gorm:"embedded;embeddedPrefix:multiplier_" json:"multiplier"`
	TickSizeFactor  int                  `json:"tickSizeFactor"`
	AmountPrecision int                  `json:"amountPrecision"`
	PricePrecision  int                  `json:"pricePrecision"`
	Status          enums_bot.Status     `json:"status"`
	Deal            BotDealModel         `gorm:"embedded;embeddedPrefix:deal_" json:"deal"`
}

func (BotModel) TableName() string {
	return "bots"
}
