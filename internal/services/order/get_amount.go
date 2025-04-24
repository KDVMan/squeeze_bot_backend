package services_order

import (
	enums_exchange "backend/internal/enums/exchange"
)

func (object *orderServiceImplementation) GetAmount() float64 {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	amount := 0.0

	for _, orderModel := range object.data {
		if orderModel.SideType == enums_exchange.SideTypeBuy &&
			orderModel.OrderType == enums_exchange.OrderTypeLimit &&
			orderModel.Status == enums_exchange.OrderStatusNew {
			amount += orderModel.Amount
		}
	}

	return amount
}
