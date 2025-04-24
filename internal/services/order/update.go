package services_order

import (
	enums_exchange "backend/internal/enums/exchange"
	models_order "backend/internal/models/order"
)

func (object *orderServiceImplementation) Update(orderModel *models_order.OrderModel) {
	orderModel.UpdateAmount()

	object.mutex.Lock()
	defer object.mutex.Unlock()

	if orderModel.Status == enums_exchange.OrderStatusCanceled ||
		orderModel.Status == enums_exchange.OrderStatusExpired ||
		(orderModel.Status == enums_exchange.OrderStatusFilled && orderModel.SideType == enums_exchange.SideTypeSell) {
		delete(object.data, orderModel.OrderID)

		if err := object.storageService().DB().Where("order_id = ?", orderModel.OrderID).Delete(&models_order.OrderModel{}).Error; err != nil {
			object.loggerService().Error().Printf("failed to delete order from DB: %v", err)
		}

		// что бы mutex не блочил т.к. там будет вызов GetAmount
		go object.userService().UpdateAvailableBalance()

		return
	}

	existingOrderModel, exists := object.data[orderModel.OrderID]
	if !exists {
		object.data[orderModel.OrderID] = orderModel

		if err := object.storageService().DB().Create(orderModel).Error; err != nil {
			object.loggerService().Error().Printf("failed to create order in DB: %v", err)
		}

		// что бы mutex не блочил т.к. там будет вызов GetAmount
		go object.userService().UpdateAvailableBalance()

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

	if err := object.storageService().DB().Save(existingOrderModel).Error; err != nil {
		object.loggerService().Error().Printf("failed to update order in DB: %v", err)
	}

	// что бы mutex не блочил т.к. там будет вызов GetAmount
	go object.userService().UpdateAvailableBalance()
}
