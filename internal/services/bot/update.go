package services_bot

import (
	"backend/internal/enums"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
)

func (object *botServiceImplementation) Update(botModel *models_bot.BotModel) error {
	if err := object.storageService().DB().Save(botModel).Error; err != nil {
		return err
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBot,
		Data:  object.LoadByHash(botModel.Hash),
	}

	return nil
}
