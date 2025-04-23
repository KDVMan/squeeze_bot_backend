package services_interface_exchange_order

import models_bot "backend/internal/models/bot"

type ExchangeOrderService interface {
	Start()
	Stop()
	AddOrder(*models_bot.BotModel, float64, float64) error
}
