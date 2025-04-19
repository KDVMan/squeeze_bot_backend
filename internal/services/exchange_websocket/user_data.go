package services_exchange_websocket

import (
	"github.com/adshao/go-binance/v2/futures"
	"strconv"
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
					for _, balance := range event.AccountUpdate.Balances {
						if balance.Asset == "USDT" {
							if amount, err := strconv.ParseFloat(balance.ChangeBalance, 64); err == nil {
								object.userService().UpdateBalance(amount)
							}
						}
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
