package models_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	"backend/internal/models"
)

type BotModel struct {
	models.DbModelWithID
	Hash           string               `gorm:"uniqueIndex:unique_bot_01;not null" json:"-"`
	Deposit        float64              `json:"deposit"`
	IsReal         bool                 `json:"isReal"`
	Symbol         string               `json:"symbol"`
	Interval       enums.Interval       `json:"interval"`
	TradeDirection enums.TradeDirection `json:"tradeDirection"`
	Window         int                  `gorm:"default:0" json:"window"`
	PrevParam      BotParamModel        `gorm:"embedded;embeddedPrefix:prev_param_" json:"prevParam"`
	CurrentParam   BotParamModel        `gorm:"embedded;embeddedPrefix:current_param_" json:"currentParam"`
	NextParam      BotParamModel        `gorm:"embedded;embeddedPrefix:next_param_" json:"nextParam"`
	Status         enums_bot.Status     `json:"status"`
}

func (BotModel) TableName() string {
	return "bots"
}
