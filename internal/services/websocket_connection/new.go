package services_websocket_connection

import (
	models_channel "backend/internal/models/channel"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_websocket_connection "backend/internal/services/websocket_connection/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"github.com/gorilla/websocket"
)

type websocketConnectionServiceImplementation struct {
	loggerService    func() services_interface_logger.LoggerService
	websocketService func() services_interface_websocket.WebsocketService
	websocket        *websocket.Conn
	broadcastChannel chan *models_channel.BroadcastChannelModel
}

func NewWebsocketConnectionService(
	loggerService func() services_interface_logger.LoggerService,
	websocketService func() services_interface_websocket.WebsocketService,
	websocket *websocket.Conn,
) services_interface_websocket_connection.WebsocketConnectionService {
	return &websocketConnectionServiceImplementation{
		loggerService:    loggerService,
		websocketService: websocketService,
		websocket:        websocket,
		broadcastChannel: make(chan *models_channel.BroadcastChannelModel, 1000),
	}
}

func (object *websocketConnectionServiceImplementation) GetBroadcastChannel() chan *models_channel.BroadcastChannelModel {
	return object.broadcastChannel
}
