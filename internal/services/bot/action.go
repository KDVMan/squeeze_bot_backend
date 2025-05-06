package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
	models_deal "backend/internal/models/deal"
)

func (object *botServiceImplementation) Action(request *models_bot.ActionRequestModel) error {
	switch request.Action {
	case enums_bot.ActionStartAll:
		return object.actionStartAll()
	case enums_bot.ActionStopAll:
		return object.actionStopAll()
	case enums_bot.ActionStopAllNotDeal:
		return object.actionStopAllNotDeal()
	case enums_bot.ActionDeleteAll:
		return object.actionDeleteAll()
	}

	return nil
}

func (object *botServiceImplementation) actionStartAll() error {
	var botsModels []models_bot.BotModel

	if err := object.storageService().DB().
		Model(&models_bot.BotModel{}).
		Where("status NOT IN ?", []enums_bot.Status{enums_bot.StatusNew, enums_bot.StatusAdd, enums_bot.StatusRun}).
		Update("status", enums_bot.StatusAdd).Error; err != nil {
		return err
	}

	if err := object.storageService().DB().
		Where("status = ?", enums_bot.StatusAdd).
		Find(&botsModels).Error; err != nil {
		return err
	}

	for _, botModel := range botsModels {
		go func(bot *models_bot.BotModel) {
			object.GetRunChannel() <- bot
		}(&botModel)
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}

func (object *botServiceImplementation) actionStopAll() error {
	var botsModels []models_bot.BotModel

	if err := object.storageService().DB().
		Where("status IN ?", []enums_bot.Status{enums_bot.StatusNew, enums_bot.StatusAdd, enums_bot.StatusRun}).
		Find(&botsModels).Error; err != nil {
		return err
	}

	for _, botModel := range botsModels {
		repositoryBot, exists := object.botRepositoryService().GetByID(botModel.ID, true)
		if exists {
			if repositoryBot.Deal.Status == enums_bot.DealStatusOpen {
				if err := object.exchangeService().AddOutMarket(repositoryBot, repositoryBot.Deal.AmountOut); err != nil {
					object.loggerService().Error().Printf("failed to add out market: %v", err)
					continue
				}
			} else if repositoryBot.OrderID != "" {
				if err := object.exchangeService().CancelLimit(repositoryBot); err != nil {
					object.loggerService().Error().Printf("failed to cancel limit: %v", err)
					continue
				}
			}

			repositoryBot.OrderID = ""
			repositoryBot.Status = enums_bot.StatusStop

			if err := object.storageService().DB().Save(&repositoryBot).Error; err != nil {
				continue
			}
		} else {
			botModel.OrderID = ""
			botModel.Status = enums_bot.StatusStop

			if err := object.storageService().DB().Save(&botModel).Error; err != nil {
				continue
			}
		}
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}

func (object *botServiceImplementation) actionStopAllNotDeal() error {
	var botsModels []models_bot.BotModel

	if err := object.storageService().DB().
		Where("status IN ?", []enums_bot.Status{enums_bot.StatusNew, enums_bot.StatusAdd, enums_bot.StatusRun}).
		Find(&botsModels).Error; err != nil {
		return err
	}

	for _, botModel := range botsModels {
		repositoryBot, exists := object.botRepositoryService().GetByID(botModel.ID, false)

		if exists && repositoryBot.Deal.Status != enums_bot.DealStatusOpen {
			if repositoryBot.OrderID != "" {
				if err := object.exchangeService().CancelLimit(repositoryBot); err != nil {
					object.loggerService().Error().Printf("failed to cancel limit: %v", err)
					continue
				}
			}

			repositoryBot.OrderID = ""
			repositoryBot.Status = enums_bot.StatusStop

			if err := object.storageService().DB().Save(&repositoryBot).Error; err != nil {
				continue
			}
		}
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}

func (object *botServiceImplementation) actionDeleteAll() error {
	var botsModels []models_bot.BotModel

	if err := object.storageService().DB().
		Where("status = ?", enums_bot.StatusStop).
		Find(&botsModels).Error; err != nil {
		return err
	}

	for _, botModel := range botsModels {
		if err := object.storageService().DB().
			Where("bot_id = ?", botModel.ID).
			Delete(&models_deal.DealModel{}).Error; err != nil {
			continue
		}

		if err := object.storageService().DB().Delete(&botModel).Error; err != nil {
			continue
		}
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}
