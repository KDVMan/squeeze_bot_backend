package services_interface_bot

import (
	models_bot "backend/internal/models/bot"
)

type BotService interface {
	Add(*models_bot.AddRequestModel) error
	AddCalculator(*models_bot.AddCalculatorRequestModel) error
	LoadAll() []*models_bot.BotModel
	LoadByID(uint) *models_bot.BotModel
	LoadByHash(string) *models_bot.BotModel
	UpdateStatus(*models_bot.UpdateStatusRequestModel) error
	RunChannel()
	GetRunChannel() chan *models_bot.BotModel
	RunDealChannel()
	GetDealChannel() chan string
	RunAddDealChannel()
	GetAddDealChannel() chan *models_bot.BotModel
	GetAmount() float64
	Update(*models_bot.BotModel) error
	UpdateParam(*models_bot.BotModel)
}
