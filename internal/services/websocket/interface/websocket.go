package services_interface_websocket

import (
	models_channel "backend/internal/models/channel"
	services_interface_websocket_connection "backend/internal/services/websocket_connection/interface"
)

type WebsocketService interface {
	Start()
	Stop()
	GetRegisterChannel() chan services_interface_websocket_connection.WebsocketConnectionService
	GetUnregisterChannel() chan services_interface_websocket_connection.WebsocketConnectionService
	GetBroadcastChannel() chan *models_channel.BroadcastChannelModel
	// GetProgressChannel() chan *models_channel.ProgressChannelModel
}
