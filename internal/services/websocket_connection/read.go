package services_websocket_connection

import (
	"errors"
	"github.com/gorilla/websocket"
)

func (object *websocketConnectionServiceImplementation) Read() {
	object.loggerService().Info().Printf("starting websocket service")

	defer func() {
		object.websocketService().GetUnregisterChannel() <- object

		if err := object.websocket.Close(); err != nil {
			object.loggerService().Error().Printf("failed to close websocket connection: %v", err)
			return
		}
	}()

	for {
		_, _, err := object.websocket.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) || errors.Is(err, websocket.ErrCloseSent) {
				return
			}

			object.loggerService().Error().Printf("failed to read websocket message: %v", err)

			return
		}
	}
}
