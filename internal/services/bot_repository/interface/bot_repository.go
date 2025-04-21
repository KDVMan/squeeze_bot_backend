package services_interface_bot_repository

import models_bot "backend/internal/models/bot"

type BotRepositoryService interface {
	Add(model *models_bot.BotModel)
	GetBySymbol(string) ([]*models_bot.BotModel, bool)
	// GetBySymbol(string) ([]*models_trade.TradeModel, bool)
	// GetAll() []*models_trade.TradeModel
}
