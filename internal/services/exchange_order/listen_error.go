package services_exchange_order

import "log"

func (object *exchangeOrderServiceImplementation) listenError() {
	for err := range object.orderPlaceService.GetReadErrorChannel() {
		log.Println("Order WS error:", err)
	}
}
