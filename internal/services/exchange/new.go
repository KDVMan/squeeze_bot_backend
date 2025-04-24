package services_exchange

import (
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_exchange_limit "backend/internal/services/exchange_limit"
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"net/http"
	"time"
)

type exchangeServiceImplementation struct {
	configService        func() services_interface_config.ConfigService
	storageService       func() services_interface_storage.StorageService
	dumpService          func() services_interface_dump.DumpService
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService
	client               *futures.Client
	listenKey            string
	stopRenewListenKey   chan struct{}
}

func NewExchangeService(
	configService func() services_interface_config.ConfigService,
	storageService func() services_interface_storage.StorageService,
	dumpService func() services_interface_dump.DumpService,
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService,
) services_interface_exchange.ExchangeService {
	client := binance.NewFuturesClient(configService().GetConfig().Binance.ApiKey, configService().GetConfig().Binance.ApiSecret)

	client.HTTPClient = &http.Client{
		Transport: &services_exchange_limit.Transport{Value: http.DefaultTransport},
	}

	return &exchangeServiceImplementation{
		storageService:       storageService,
		exchangeLimitService: exchangeLimitService,
		dumpService:          dumpService,
		client:               client,
		listenKey:            "",
		stopRenewListenKey:   nil,
	}
}

func (object *exchangeServiceImplementation) getOrderID(id uint) string {
	return fmt.Sprintf("order-%d-%d", id, time.Now().UnixNano())
}
