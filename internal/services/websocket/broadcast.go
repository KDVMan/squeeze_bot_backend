package services_websocket

import (
	models_channel "backend/internal/models/channel"
)

func (object *websocketServiceImplementation) broadcast(broadcastModel *models_channel.BroadcastChannelModel) {
	object.lock.Lock()
	defer object.lock.Unlock()

	for connection := range object.connections {
		select {
		case connection.GetBroadcastChannel() <- broadcastModel:
			// object.loggerService().Info().Printf("sent message: %v", broadcastModel.Event)
		default:
			object.loggerService().Error().Printf("failed to send message: %v", broadcastModel)
		}
	}
}
