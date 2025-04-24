package services_user

import (
	models_user "backend/internal/models/user"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_order "backend/internal/services/order/interface"
	services_interface_user "backend/internal/services/user/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
	"sync"
)

type userServiceImplementation struct {
	loggerService    func() services_interface_logger.LoggerService
	storageService   func() services_interface_storage.StorageService
	websocketService func() services_interface_websocket.WebsocketService
	dumpService      func() services_interface_dump.DumpService
	exchangeService  func() services_interface_exchange.ExchangeService
	orderService     func() services_interface_order.OrderService
	userModel        *models_user.UserModel
	userMutex        *sync.Mutex
}

func NewUserService(
	loggerService func() services_interface_logger.LoggerService,
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	dumpService func() services_interface_dump.DumpService,
	exchangeService func() services_interface_exchange.ExchangeService,
	orderService func() services_interface_order.OrderService,
) services_interface_user.UserService {
	return &userServiceImplementation{
		loggerService:    loggerService,
		storageService:   storageService,
		websocketService: websocketService,
		dumpService:      dumpService,
		exchangeService:  exchangeService,
		orderService:     orderService,
		userModel:        &models_user.UserModel{},
		userMutex:        &sync.Mutex{},
	}
}
