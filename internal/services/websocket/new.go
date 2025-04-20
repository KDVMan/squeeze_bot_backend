package services_websocket

import (
	models_channel "backend/internal/models/channel"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_user "backend/internal/services/user/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_websocket_connection "backend/internal/services/websocket_connection/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
)

type websocketServiceImplementation struct {
	loggerService        func() services_interface_logger.LoggerService
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService
	symbolService        func() services_interface_symbol.SymbolService
	userService          func() services_interface_user.UserService
	botService           func() services_interface_bot.BotService
	connections          map[services_interface_websocket_connection.WebsocketConnectionService]bool
	registerChannel      chan services_interface_websocket_connection.WebsocketConnectionService
	unregisterChannel    chan services_interface_websocket_connection.WebsocketConnectionService
	broadcastChannel     chan *models_channel.BroadcastChannelModel
	// progressChannel      chan *models_channel.ProgressChannelModel
	lock sync.Mutex
}

func NewWebsocketService(
	loggerService func() services_interface_logger.LoggerService,
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService,
	symbolService func() services_interface_symbol.SymbolService,
	userService func() services_interface_user.UserService,
	botService func() services_interface_bot.BotService,
) services_interface_websocket.WebsocketService {
	return &websocketServiceImplementation{
		loggerService:        loggerService,
		exchangeLimitService: exchangeLimitService,
		symbolService:        symbolService,
		userService:          userService,
		botService:           botService,
		connections:          make(map[services_interface_websocket_connection.WebsocketConnectionService]bool),
		registerChannel:      make(chan services_interface_websocket_connection.WebsocketConnectionService, 1000),
		unregisterChannel:    make(chan services_interface_websocket_connection.WebsocketConnectionService, 1000),
		broadcastChannel:     make(chan *models_channel.BroadcastChannelModel, 1000),
		// progressChannel:      make(chan *models_channel.ProgressChannelModel, 1000),
	}
}

func (object *websocketServiceImplementation) GetRegisterChannel() chan services_interface_websocket_connection.WebsocketConnectionService {
	return object.registerChannel
}

func (object *websocketServiceImplementation) GetUnregisterChannel() chan services_interface_websocket_connection.WebsocketConnectionService {
	return object.unregisterChannel
}

func (object *websocketServiceImplementation) GetBroadcastChannel() chan *models_channel.BroadcastChannelModel {
	return object.broadcastChannel
}

// func (object *websocketServiceImplementation) GetProgressChannel() chan *models_channel.ProgressChannelModel {
// 	return object.progressChannel
// }
