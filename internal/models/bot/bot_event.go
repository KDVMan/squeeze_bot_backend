package models_bot

import (
	enums_bot "backend/internal/enums/bot"
	models_order "backend/internal/models/order"
)

type BotEventModel struct {
	BotID  uint
	Symbol string
	Type   enums_bot.BotEvent
	Order  *models_order.OrderModel
}
