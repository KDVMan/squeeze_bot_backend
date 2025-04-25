package models_deal

import (
	"backend/internal/enums"
	"backend/internal/models"
)

type DealModel struct {
	models.DbModelWithID
	BotID          uint                 `gorm:"not null;index" json:"botID"`
	IsReal         bool                 `json:"isReal"`
	TradeDirection enums.TradeDirection `json:"tradeDirection"`
	TimeIn         int64                `json:"timeIn"`
	TimeOut        int64                `json:"timeOut"`
	PriceIn        float64              `json:"priceIn"`
	AmountIn       float64              `json:"amountIn"`
	PriceOut       float64              `json:"priceOut"`
	AmountOut      float64              `json:"amountOut"`
	IsStopTime     bool                 `json:"isStopTime"`
	IsStopPercent  bool                 `json:"isStopPercent"`
	ProfitPercent  float64              `json:"profitPercent"`
	Param          DealParamModel       `gorm:"embedded;embeddedPrefix:param_" json:"param"`
}

func (DealModel) TableName() string {
	return "deals"
}
