package services_interface_bot

import (
	models_bot "backend/internal/models/bot"
)

type BotService interface {
	Add(*models_bot.AddRequestModel) error
	LoadAll() []*models_bot.BotModel
	LoadByHash(string) *models_bot.BotModel
	Status(*models_bot.StatusRequestModel) error
	RunChannel()
	GetRunChannel() chan *models_bot.BotModel
	RunDealChannel()
	GetDealChannel() chan string
	RunAddDealChannel()
	GetAddDealChannel() chan *models_bot.BotModel
}
