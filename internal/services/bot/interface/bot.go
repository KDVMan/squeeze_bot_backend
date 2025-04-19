package services_interface_bot

import (
	models_bot "backend/internal/models/bot"
)

type BotService interface {
	Start(*models_bot.StartRequestModel) error
}
