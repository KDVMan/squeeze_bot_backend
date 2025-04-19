package services_router

import (
	services_websocket_connection "backend/internal/services/websocket_connection"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (object *routerServiceImplementation) websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		object.loggerService().Error().Printf("failed to upgrade websocket: %v", err)
		return
	}

	connection := services_websocket_connection.NewWebsocketConnectionService(
		object.loggerService,
		object.websocketService,
		conn,
	)

	object.websocketService().GetRegisterChannel() <- connection

	go connection.Write()
	go connection.Read()
}
