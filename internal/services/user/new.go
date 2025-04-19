package services_user

import (
	models_user "backend/internal/models/user"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_user "backend/internal/services/user/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
	"sync"
)

type userServiceImplementation struct {
	storageService   func() services_interface_storage.StorageService
	websocketService func() services_interface_websocket.WebsocketService
	exchangeService  func() services_interface_exchange.ExchangeService
	userModel        *models_user.UserModel
	mutex            *sync.Mutex
}

func NewUserService(
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	exchangeService func() services_interface_exchange.ExchangeService,
) services_interface_user.UserService {
	return &userServiceImplementation{
		storageService:   storageService,
		websocketService: websocketService,
		exchangeService:  exchangeService,
		userModel:        models_user.LoadDefault(),
		mutex:            &sync.Mutex{},
	}
}
