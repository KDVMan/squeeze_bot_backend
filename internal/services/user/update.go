package services_user

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
	models_user "backend/internal/models/user"
)

func (object *userServiceImplementation) Update(userModel *models_user.UserModel, broadcast bool) {
	object.userModel = userModel

	if broadcast {
		broadcastModel := models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventUser,
			Data:  object.userModel,
		}

		object.websocketService().GetBroadcastChannel() <- &broadcastModel
	}
}
