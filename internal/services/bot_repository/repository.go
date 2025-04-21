package services_bot_repository

import (
	models_bot "backend/internal/models/bot"
)

func (object *botRepositoryServiceImplementation) Add(botModel *models_bot.BotModel) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	object.data[botModel.Symbol] = append(object.data[botModel.Symbol], botModel)
}

func (object *botRepositoryServiceImplementation) GetBySymbol(symbol string) ([]*models_bot.BotModel, bool) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	botsModels, exists := object.data[symbol]

	return botsModels, exists
}

// func (object *botRepositoryServiceImplementation) GetAll() []*models_trade.TradeModel {
// 	var tradesModels []*models_trade.TradeModel
//
// 	object.mutex.Lock()
// 	defer object.mutex.Unlock()
//
// 	if len(object.data) == 0 {
// 		return nil
// 	}
//
// 	for _, trades := range object.data {
// 		tradesModels = append(tradesModels, trades...)
// 	}
//
// 	return tradesModels
// }
