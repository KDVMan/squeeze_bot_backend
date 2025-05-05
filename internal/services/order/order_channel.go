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

		botsModels, exists := object.botRepositoryService().GetBySymbol(orderModel.Symbol)
		if !exists {
			continue
		}

		for _, botModel := range botsModels {
			if botModel.OrderID != orderModel.OrderID {
				continue
			}

			if orderModel.PositionType == enums_exchange.PositionTypeLong {
				if orderModel.SideType == enums_exchange.SideTypeBuy &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusNew &&
					orderModel.Status == enums_exchange.OrderStatusNew {
					object.statusOpenLimit(orderModel, botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeBuy &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusAmendment &&
					orderModel.Status == enums_exchange.OrderStatusNew {
					object.statusOpenLimit(orderModel, botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeBuy &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusCanceled &&
					orderModel.Status == enums_exchange.OrderStatusCanceled {
					object.cancelLimit(botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeBuy &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusTrade &&
					orderModel.Status == enums_exchange.OrderStatusFilled {
					object.openLimit(orderModel, botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeSell &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusTrade &&
					orderModel.Status == enums_exchange.OrderStatusFilled {
					object.closeLimit(orderModel, botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeSell &&
					orderModel.OrderType == enums_exchange.OrderTypeMarket &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusTrade &&
					orderModel.Status == enums_exchange.OrderStatusFilled {
					object.closeLimit(orderModel, botModel)
				}
			} else if orderModel.PositionType == enums_exchange.PositionTypeShort {
				if orderModel.SideType == enums_exchange.SideTypeSell &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusNew &&
					orderModel.Status == enums_exchange.OrderStatusNew {
					object.statusOpenLimit(orderModel, botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeSell &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusAmendment &&
					orderModel.Status == enums_exchange.OrderStatusNew {
					object.statusOpenLimit(orderModel, botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeSell &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusCanceled &&
					orderModel.Status == enums_exchange.OrderStatusCanceled {
					object.cancelLimit(botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeSell &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusTrade &&
					orderModel.Status == enums_exchange.OrderStatusFilled {
					object.openLimit(orderModel, botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeBuy &&
					orderModel.OrderType == enums_exchange.OrderTypeLimit &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusTrade &&
					orderModel.Status == enums_exchange.OrderStatusFilled {
					object.closeLimit(orderModel, botModel)
				} else if orderModel.SideType == enums_exchange.SideTypeBuy &&
					orderModel.OrderType == enums_exchange.OrderTypeMarket &&
					orderModel.ExecutionStatus == enums_exchange.OrderExecutionStatusTrade &&
					orderModel.Status == enums_exchange.OrderStatusFilled {
					object.closeLimit(orderModel, botModel)
				}
			}
		}
	}
}

func (object *orderServiceImplementation) statusOpenLimit(orderModel *models_order.OrderModel, botModel *models_bot.BotModel) {
	object.botService().GetBotEventChannel() <- &models_bot.BotEventModel{
		BotID:  botModel.ID,
		Symbol: botModel.Symbol,
		Type:   enums_bot.BotEventStatusOpenLimit,
		Order:  orderModel,
	}
}

func (object *orderServiceImplementation) cancelLimit(botModel *models_bot.BotModel) {
	object.botService().GetBotEventChannel() <- &models_bot.BotEventModel{
		BotID:  botModel.ID,
		Symbol: botModel.Symbol,
		Type:   enums_bot.BotEventCancelLimit,
	}
}

func (object *orderServiceImplementation) openLimit(orderModel *models_order.OrderModel, botModel *models_bot.BotModel) {
	object.botService().GetBotEventChannel() <- &models_bot.BotEventModel{
		BotID:  botModel.ID,
		Symbol: botModel.Symbol,
		Type:   enums_bot.BotEventOpenLimit,
		Order:  orderModel,
	}
}

func (object *orderServiceImplementation) closeLimit(orderModel *models_order.OrderModel, botModel *models_bot.BotModel) {
	object.botService().GetBotEventChannel() <- &models_bot.BotEventModel{
		BotID:  botModel.ID,
		Symbol: botModel.Symbol,
		Type:   enums_bot.BotEventCloseLimit,
		Order:  orderModel,
	}
}

func (object *orderServiceImplementation) GetOrderChannel() chan *models_order.OrderModel {
	return object.orderChannel
}
