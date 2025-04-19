package services_websocket

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
)

func (object *websocketServiceImplementation) broadcastSymbols() {
	symbols, err := object.symbolService().LoadAll()

	if err != nil {
		object.loggerService().Error().Printf("failed to load symbols: %v", err)
		return
	}

	object.broadcastChannel <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventSymbolList,
		Data:  symbols,
	}
}
