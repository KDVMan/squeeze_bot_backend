package services_bot

import (
	models_bot "backend/internal/models/bot"
	"errors"
	"gorm.io/gorm"
)

func (object *botServiceImplementation) LoadAll() []*models_bot.BotModel {
	var botsModels []*models_bot.BotModel

	initModel, err := object.initService().Load()
	if err != nil {
		object.loggerService().Error().Printf("failed to load init: %v", err)
		return []*models_bot.BotModel{}
	}

	sortColumn := initModel.BotSortColumn.DB()
	sortDirection := initModel.BotSortDirection.String()

	if err = object.storageService().DB().
		Order(sortColumn + " " + sortDirection).
		Find(&botsModels).Error; err != nil {

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			object.loggerService().Error().Printf("failed to load bots: %v", err)
		}

		return []*models_bot.BotModel{}
	}

	repoBots := object.botRepositoryService().GetAll()

	for i, dbBot := range botsModels {
		for _, repoBot := range repoBots {
			if dbBot.ID == repoBot.ID {
				botsModels[i] = repoBot
				break
			}
		}
	}

	return botsModels
}
