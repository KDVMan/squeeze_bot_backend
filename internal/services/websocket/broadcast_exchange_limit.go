package services_websocket

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
)

func (object *websocketServiceImplementation) broadcastExchangeLimits() {
	limits, err := object.exchangeLimitService().Load()

	if err != nil {
		object.loggerService().Error().Printf("failed to load exchange limits: %v", err)
		return
	}

	object.broadcastChannel <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventExchangeLimits,
		Data:  limits,
	}
}
