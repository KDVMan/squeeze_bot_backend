package services_interface_bot

import (
	models_bot "backend/internal/models/bot"
)

type BotService interface {
	Start(*models_bot.StartRequestModel) error
	Load() []*models_bot.BotModel
	Status(*models_bot.StatusRequestModel) error
}
