package services_exchange_order

import (
	services_interface_exchange_order "backend/internal/services/exchange_order/interface"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
	"sync"
	"time"
)

type exchangeOrderServiceImplementation struct {
	loggerService     func() services_interface_logger.LoggerService
	configService     func() services_interface_config.ConfigService
	dumpService       func() services_interface_dump.DumpService
	apiKey            string
	apiSecret         string
	orderPlaceService *futures.OrderPlaceWsService
	mutex             *sync.Mutex
}

func NewExchangeOrderService(
	loggerService func() services_interface_logger.LoggerService,
	configService func() services_interface_config.ConfigService,
	dumpService func() services_interface_dump.DumpService,
) services_interface_exchange_order.ExchangeOrderService {
	service := &exchangeOrderServiceImplementation{
		loggerService: loggerService,
		configService: configService,
		dumpService:   dumpService,
		apiKey:        configService().GetConfig().Binance.ApiKey,
		apiSecret:     configService().GetConfig().Binance.ApiSecret,
		mutex:         &sync.Mutex{},
	}

	orderPlaceService, err := futures.NewOrderPlaceWsService(service.apiKey, service.apiSecret)
	if err != nil {
		loggerService().Error().Panicf("failed to create order place service: %v", err)
	}

	service.orderPlaceService = orderPlaceService

	return service
}

func (object *exchangeOrderServiceImplementation) Start() {
	object.loggerService().Info().Printf("starting order websocket service...")

	go object.listenResponse()
	go object.listenError()
}

func (object *exchangeOrderServiceImplementation) Stop() {
	object.loggerService().Info().Println("stopping order websocket service...")

	object.orderPlaceService.ReceiveAllDataBeforeStop(10 * time.Second)
	object.orderPlaceService = nil
}

func (object *exchangeOrderServiceImplementation) getOrderID(id uint) string {
	return fmt.Sprintf("order-%d-%d", id, time.Now().UnixNano())
}
