package services_interface_bot_repository

import models_bot "backend/internal/models/bot"

type BotRepositoryService interface {
	Add(*models_bot.BotModel)
	GetByID(uint, bool) (*models_bot.BotModel, bool)
	GetBySymbol(string) ([]*models_bot.BotModel, bool)
	GetBySymbolAndID(string, uint) (*models_bot.BotModel, bool)
	GetAll() []*models_bot.BotModel
	Remove(string, uint)
}
