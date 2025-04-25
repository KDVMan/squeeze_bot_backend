package services_user

import (
	models_user "backend/internal/models/user"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_user "backend/internal/services/user/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
)

type userServiceImplementation struct {
	loggerService    func() services_interface_logger.LoggerService
	websocketService func() services_interface_websocket.WebsocketService
	dumpService      func() services_interface_dump.DumpService
	exchangeService  func() services_interface_exchange.ExchangeService
	botService       func() services_interface_bot.BotService
	userModel        *models_user.UserModel
	userMutex        *sync.Mutex
}

func NewUserService(
	loggerService func() services_interface_logger.LoggerService,
	websocketService func() services_interface_websocket.WebsocketService,
	dumpService func() services_interface_dump.DumpService,
	exchangeService func() services_interface_exchange.ExchangeService,
	botService func() services_interface_bot.BotService,
) services_interface_user.UserService {
	return &userServiceImplementation{
		loggerService:    loggerService,
		websocketService: websocketService,
		dumpService:      dumpService,
		exchangeService:  exchangeService,
		botService:       botService,
		userMutex:        &sync.Mutex{},
	}
}
