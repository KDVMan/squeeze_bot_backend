package services_user

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
)

func (object *userServiceImplementation) UpdateAvailableBalance() {
	object.userMutex.Lock()
	defer object.userMutex.Unlock()

	object.userModel.AvailableBalance = object.userModel.Balance - object.orderService().GetAmount()

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventUser,
		Data:  object.userModel,
	}
}
