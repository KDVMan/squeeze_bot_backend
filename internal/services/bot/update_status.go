package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
	models_deal "backend/internal/models/deal"
)

func (object *botServiceImplementation) UpdateStatus(request *models_bot.UpdateStatusRequestModel) error {
	var storageBotModel models_bot.BotModel

	err := object.storageService().DB().First(&storageBotModel, request.ID).Error
	if err != nil {
		return err
	}

	if request.Status == enums_bot.StatusAdd {
		object.GetRunChannel() <- &storageBotModel
	} else if request.Status == enums_bot.StatusStop {
		botModel, exists := object.botRepositoryService().GetByID(storageBotModel.ID, true)
		if exists {
			if botModel.Deal.Status == enums_bot.DealStatusOpen {
				if err = object.exchangeService().AddOutMarket(botModel, botModel.Deal.AmountOut); err != nil {
					object.loggerService().Error().Printf("failed to add out market: %v", err)
					return err
				}
			} else if botModel.OrderID != "" {
				if err = object.exchangeService().CancelLimit(botModel); err != nil {
					object.loggerService().Error().Printf("failed to cancel limit: %v", err)
					return err
				}
			}
		} else {
			botModel = &storageBotModel
		}

		botModel.OrderID = ""
		botModel.Status = enums_bot.StatusStop

		if err = object.storageService().DB().Save(&botModel).Error; err != nil {
			return err
		}

		object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventBot,
			Data:  botModel,
		}
	} else if request.Status == enums_bot.StatusDelete {
		if err = object.storageService().DB().
			Where("bot_id = ?", storageBotModel.ID).
			Delete(&models_deal.DealModel{}).Error; err != nil {
			return err
		}

		if err = object.storageService().DB().Delete(&storageBotModel).Error; err != nil {
			return err
		}

		object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventBotList,
			Data:  object.LoadAll(),
		}
	}

	return nil
}
