package services_user

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
)

func (object *userServiceImplementation) UpdateBalance(balance float64) {
	object.userMutex.Lock()
	defer object.userMutex.Unlock()

	object.userModel.Balance = balance
	object.userModel.AvailableBalance = balance - object.botService().GetAmount()

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventUser,
		Data:  object.userModel,
	}
}
