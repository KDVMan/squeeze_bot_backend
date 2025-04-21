package services_exchange_binance

import (
	services_interface_exchange_binance "backend/pkg/services/exchange_binance/interface"
)

type exchangeBinanceServiceImplementation struct {
}

func NewExchangeBinanceService() services_interface_exchange_binance.ExchangeBinanceService {
	return &exchangeBinanceServiceImplementation{}
}
