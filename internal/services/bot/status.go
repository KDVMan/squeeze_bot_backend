package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
)

func (object *botServiceImplementation) Status(request *models_bot.StatusRequestModel) error {
	var botModel models_bot.BotModel

	err := object.storageService().DB().First(&botModel, request.ID).Error
	if err != nil {
		return err
	}

	botModel.Status = request.Status

	if err = object.storageService().DB().Save(&botModel).Error; err != nil {
		return err
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBot,
		Data:  object.LoadByHash(botModel.Hash),
	}

	if botModel.Status == enums_bot.StatusAdd {
		object.GetRunChannel() <- &botModel
	}

	return nil
}
