package services_order

import (
	enums_exchange "backend/internal/enums/exchange"
	models_order "backend/internal/models/order"
)

func (object *orderServiceImplementation) Update(orderModel *models_order.OrderModel) {
	orderModel.UpdateAmount()

	object.mutex.Lock()

	defer func() {
		// что бы mutex не блочил т.к. там будет вызов GetAmount
		go object.userService().UpdateAvailableBalance()

		object.mutex.Unlock()
	}()

	if object.isOrderClosed(orderModel) {
		delete(object.data, orderModel.OrderID)

		return
	}

	existingOrderModel, exists := object.data[orderModel.OrderID]
	if !exists {
		object.data[orderModel.OrderID] = orderModel

		return
	}

	existingOrderModel.SideType = orderModel.SideType
	existingOrderModel.OrderType = orderModel.OrderType
	existingOrderModel.ExecutionStatus = orderModel.ExecutionStatus
	existingOrderModel.Status = orderModel.Status
	existingOrderModel.OriginalPrice = orderModel.OriginalPrice
	existingOrderModel.AveragePrice = orderModel.AveragePrice
	existingOrderModel.OriginalQuantity = orderModel.OriginalQuantity
	existingOrderModel.FilledQuantity = orderModel.FilledQuantity
	existingOrderModel.Commission = orderModel.Commission
	existingOrderModel.Amount = orderModel.Amount
}

func (object *orderServiceImplementation) isOrderClosed(orderModel *models_order.OrderModel) bool {
	if orderModel.Status == enums_exchange.OrderStatusCanceled {
		return true
	} else if orderModel.Status == enums_exchange.OrderStatusExpired {
		return true
	} else if orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusTrade && orderModel.Status == enums_exchange.OrderStatusFilled {
		return true
	}

	return false
}
