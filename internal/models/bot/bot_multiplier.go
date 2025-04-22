package models_bot

import (
	"backend/internal/enums"
)

type BotMultiplierModel struct {
	Value      float64
	MinKeyName enums.Bind
	MaxKeyName enums.Bind
}

func GetMultiplier(tradeDirection enums.TradeDirection) BotMultiplierModel {
	if tradeDirection == enums.TradeDirectionShort {
		return BotMultiplierModel{
			Value:      -1,
			MinKeyName: enums.BindHigh,
			MaxKeyName: enums.BindLow,
		}
	}

	return BotMultiplierModel{
		Value:      1,
		MinKeyName: enums.BindLow,
		MaxKeyName: enums.BindHigh,
	}
}
