package services_bot_repository

import (
	models_bot "backend/internal/models/bot"
)

func (object *botRepositoryServiceImplementation) Add(botModel *models_bot.BotModel) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	object.data[botModel.Symbol] = append(object.data[botModel.Symbol], botModel)
}

func (object *botRepositoryServiceImplementation) GetByID(ID uint, remove bool) (*models_bot.BotModel, bool) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	for symbol, botsModels := range object.data {
		for i, botModel := range botsModels {
			if botModel.ID == ID {
				if remove {
					object.data[symbol] = append(botsModels[:i], botsModels[i+1:]...)

					if len(object.data[symbol]) == 0 {
						delete(object.data, symbol)
					}
				}

				return botModel, true
			}
		}
	}

	return nil, false
}

func (object *botRepositoryServiceImplementation) GetBySymbol(symbol string) ([]*models_bot.BotModel, bool) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	botsModels, exists := object.data[symbol]

	return botsModels, exists
}

func (object *botRepositoryServiceImplementation) GetAll() []*models_bot.BotModel {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	if len(object.data) == 0 {
		return nil
	}

	var botsModels []*models_bot.BotModel

	for _, botModel := range object.data {
		botsModels = append(botsModels, botModel...)
	}

	return botsModels
}

func (object *botRepositoryServiceImplementation) Remove(symbol string, orderID uint) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	botsModels, exists := object.data[symbol]
	if !exists {
		return
	}

	for i, botModel := range botsModels {
		if botModel.ID == orderID {
			object.data[symbol] = append(botsModels[:i], botsModels[i+1:]...)

			if len(object.data[symbol]) == 0 {
				delete(object.data, symbol)
			}

			return
		}
	}

	return
}
