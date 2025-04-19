package services_websocket_connection

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

func (object *websocketConnectionServiceImplementation) Write() {
	for {
		broadcastModel, ok := <-object.broadcastChannel

		if !ok {
			if err := object.websocket.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
				// object.loggerService().Error().Printf("failed to write close message: %v", err)
				return
			}

			return
		}

		message, err := json.Marshal(broadcastModel)

		if err != nil {
			object.loggerService().Error().Printf("failed to marshal message: %v", err)
			continue
		}

		if err = object.websocket.WriteMessage(websocket.TextMessage, message); err != nil {
			object.loggerService().Error().Printf("failed to write text message: %v", err)
			return
		}
	}
}
