package services_order

import (
	enums_bot "backend/internal/enums/bot"
	enums_exchange "backend/internal/enums/exchange"
	models_bot "backend/internal/models/bot"
	models_order "backend/internal/models/order"
)

func (object *orderServiceImplementation) RunOrderChannel() {
	for orderModel := range object.orderChannel {
		object.Update(orderModel)

		object.dumpService().Dump(orderModel)

		botsModels, exists := object.botRepositoryService().GetBySymbol(orderModel.Symbol)
		if !exists {
			continue
		}

		for _, botModel := range botsModels {
			if botModel.OrderID != orderModel.OrderID {
				continue
			}

			if orderModel.SideType == enums_exchange.SideTypeBuy &&
				orderModel.OrderType == enums_exchange.OrderTypeLimit &&
				orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusNew &&
				orderModel.Status == enums_exchange.OrderStatusNew {
				botModel.Deal.Status = enums_bot.DealStatusOpenLimit
			} else if orderModel.SideType == enums_exchange.SideTypeBuy &&
				orderModel.OrderType == enums_exchange.OrderTypeLimit &&
				orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusAmendment &&
				orderModel.Status == enums_exchange.OrderStatusNew {
				botModel.Deal.Status = enums_bot.DealStatusOpenLimit
			} else if orderModel.SideType == enums_exchange.SideTypeBuy &&
				orderModel.OrderType == enums_exchange.OrderTypeLimit &&
				orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusCanceled &&
				orderModel.Status == enums_exchange.OrderStatusCanceled {
				botModel.OrderID = ""
				botModel.Deal = models_bot.BotDealModel{}
			}
		}
	}
}

func (object *orderServiceImplementation) GetOrderChannel() chan *models_order.OrderModel {
	return object.orderChannel
}
