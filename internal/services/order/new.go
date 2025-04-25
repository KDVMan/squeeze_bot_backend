package services_order

import (
	models_order "backend/internal/models/order"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_bot_repository "backend/internal/services/bot_repository/interface"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_order "backend/internal/services/order/interface"
	services_interface_user "backend/internal/services/user/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
)

type orderServiceImplementation struct {
	loggerService        func() services_interface_logger.LoggerService
	dumpService          func() services_interface_dump.DumpService
	websocketService     func() services_interface_websocket.WebsocketService
	exchangeService      func() services_interface_exchange.ExchangeService
	botService           func() services_interface_bot.BotService
	botRepositoryService func() services_interface_bot_repository.BotRepositoryService
	userService          func() services_interface_user.UserService
	data                 map[string]*models_order.OrderModel
	mutex                *sync.Mutex
	orderChannel         chan *models_order.OrderModel
}

func NewOrderService(
	loggerService func() services_interface_logger.LoggerService,
	dumpService func() services_interface_dump.DumpService,
	websocketService func() services_interface_websocket.WebsocketService,
	exchangeService func() services_interface_exchange.ExchangeService,
	botService func() services_interface_bot.BotService,
	botRepositoryService func() services_interface_bot_repository.BotRepositoryService,
	userService func() services_interface_user.UserService,
) services_interface_order.OrderService {
	return &orderServiceImplementation{
		loggerService:        loggerService,
		dumpService:          dumpService,
		websocketService:     websocketService,
		exchangeService:      exchangeService,
		botService:           botService,
		botRepositoryService: botRepositoryService,
		userService:          userService,
		data:                 make(map[string]*models_order.OrderModel),
		mutex:                &sync.Mutex{},
		orderChannel:         make(chan *models_order.OrderModel, 1000000),
	}
}
