package services_balance

import (
	services_interface_balance "backend/internal/services/balance/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
)

type balanceServiceImplementation struct {
	loggerService func() services_interface_logger.LoggerService
	balance       float64
	data          map[uint]float64
	mutex         *sync.Mutex
}

func NewBalanceService(
	loggerService func() services_interface_logger.LoggerService,
) services_interface_balance.BalanceService {
	return &balanceServiceImplementation{
		loggerService: loggerService,
		balance:       0,
		data:          make(map[uint]float64),
		mutex:         &sync.Mutex{},
	}
}

func (object *balanceServiceImplementation) InitBalance(balance float64) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	object.balance = balance
}
