package services_user

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
)

func (object *userServiceImplementation) Broadcast() {
	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventUser,
		Data:  object.userModel,
	}
}
