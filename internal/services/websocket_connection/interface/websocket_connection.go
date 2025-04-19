package services_interface_websocket_connection

import models_channel "backend/internal/models/channel"

type WebsocketConnectionService interface {
	Read()
	Write()
	GetBroadcastChannel() chan *models_channel.BroadcastChannelModel
}
