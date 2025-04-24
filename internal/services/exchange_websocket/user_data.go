package services_exchange_websocket

import (
	enums_exchange "backend/internal/enums/exchange"
	models_order "backend/internal/models/order"
	services_helper "backend/pkg/services/helper"
	"github.com/adshao/go-binance/v2/futures"
	"time"
)

func (object *exchangeWebsocketServiceImplementation) userData() {
	for {
		var stopChannel chan struct{}

		listenKey, err := object.exchangeService().GetListenKey()
		if err != nil {
			object.loggerService().Error().Printf("failed to get listen key: %v", err)
			return
		}

		doneChannel, stopChannel, err := futures.WsUserDataServe(
			listenKey,
			func(event *futures.WsUserDataEvent) {
				if event.Event == futures.UserDataEventTypeAccountUpdate {
					object.dumpService().Dump(event.AccountUpdate)

					for _, balance := range event.AccountUpdate.Balances {
						if balance.Asset == "USDT" {
							object.userService().UpdateBalance(services_helper.MustConvertStringToFloat64(balance.Balance))
						}
					}
				} else if event.Event == futures.UserDataEventTypeOrderTradeUpdate {
					object.dumpService().Dump(event.OrderTradeUpdate)

					object.orderService().GetOrderChannel() <- &models_order.OrderModel{
						OrderID:          event.OrderTradeUpdate.ClientOrderID,
						Symbol:           event.OrderTradeUpdate.Symbol,
						SideType:         enums_exchange.SideType(event.OrderTradeUpdate.Side),
						OrderType:        enums_exchange.OrderType(event.OrderTradeUpdate.Type),
						ExecutionStatus:  enums_exchange.OrderExecutionStatus(event.OrderTradeUpdate.ExecutionType),
						Status:           enums_exchange.OrderStatus(event.OrderTradeUpdate.Status),
						OriginalPrice:    services_helper.MustConvertStringToFloat64(event.OrderTradeUpdate.OriginalPrice),
						AveragePrice:     services_helper.MustConvertStringToFloat64(event.OrderTradeUpdate.AveragePrice),
						OriginalQuantity: services_helper.MustConvertStringToFloat64(event.OrderTradeUpdate.OriginalQty),
						FilledQuantity:   services_helper.MustConvertStringToFloat64(event.OrderTradeUpdate.AccumulatedFilledQty),
						Commission:       services_helper.MustConvertStringToFloat64(event.OrderTradeUpdate.Commission),
					}
				}
			},
			func(err error) {
				object.loggerService().Error().Printf("websocket error: %v", err)
				stopChannel <- struct{}{} // корректно закрываем текущее соединение
			},
		)
		if err != nil {
			object.loggerService().Error().Printf("failed to start websocket: %v", err)
			time.Sleep(object.reconnectDelay)
			continue
		}

		object.loggerService().Info().Println("user data - started")

		select {
		case <-doneChannel:
			object.loggerService().Info().Printf("reconnect")
			time.Sleep(object.reconnectDelay)
			continue
		case <-object.doneChannel:
			stopChannel <- struct{}{}
			return
		}
	}
}
