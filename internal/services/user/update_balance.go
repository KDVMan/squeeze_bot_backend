package services_user

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
)

func (object *userServiceImplementation) UpdateBalance(balance float64) {
	object.userMutex.Lock()
	defer object.userMutex.Unlock()

	object.userModel.Balance = balance
	object.userModel.AvailableBalance = balance - object.orderService().GetAmount()

	if object.storageService().DB().Save(object.userModel).Error != nil {
		object.loggerService().Error().Printf("failed to update user balance")
		return
	}

	object.dumpService().Dump(object.userModel)

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventUser,
		Data:  object.userModel,
	}
}
