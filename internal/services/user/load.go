package services_user

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
)

func (object *userServiceImplementation) Load() {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	totalBalance, availableBalance, err := object.exchangeService().UserBalance()
	if err != nil {
		return
	}

	hedge, err := object.exchangeService().UserHedge()
	if err != nil {
		return
	}

	object.userModel.TotalBalance = totalBalance
	object.userModel.AvailableBalance = availableBalance
	object.userModel.Hedge = hedge

	broadcastModel := models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventUser,
		Data:  object.userModel,
	}

	object.websocketService().GetBroadcastChannel() <- &broadcastModel
}
