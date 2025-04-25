package services_user

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
	"log"
)

func (object *userServiceImplementation) UpdateAvailableBalance() {
	object.userMutex.Lock()
	defer object.userMutex.Unlock()

	object.userModel.AvailableBalance = object.userModel.Balance - object.botService().GetAmount()

	log.Println("AVAILABLE BALANCE", object.userModel.AvailableBalance)

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventUser,
		Data:  object.userModel,
	}
}
