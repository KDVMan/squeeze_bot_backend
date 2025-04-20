package services_websocket

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
)

func (object *websocketServiceImplementation) Start() {
	object.loggerService().Info().Printf("starting websocket service")

	for {
		select {
		case connection := <-object.registerChannel:
			object.lock.Lock()
			object.connections[connection] = true
			object.lock.Unlock()

			object.loggerService().Info().Printf("registered websocket connection")

			go object.broadcastSymbols()
			go object.broadcastExchangeLimits()
			go object.userService().Load()

			object.broadcastChannel <- &models_channel.BroadcastChannelModel{
				Event: enums.WebsocketEventBot,
				Data:  object.botService().Load(),
			}
		case connection := <-object.unregisterChannel:
			object.lock.Lock()

			if _, ok := object.connections[connection]; ok {
				delete(object.connections, connection)
				close(connection.GetBroadcastChannel())

				object.loggerService().Info().Printf("unregistered websocket connection")
			}

			object.lock.Unlock()
		case data := <-object.broadcastChannel:
			object.broadcast(data)
			// case data := <-object.progressChannel:
			// 	object.broadcast(&models_channel.BroadcastChannelModel{
			// 		Event: data.Event,
			// 		Data:  data,
			// 	})
		}
	}
}
