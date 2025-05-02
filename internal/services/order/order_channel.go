package services_order

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	enums_exchange "backend/internal/enums/exchange"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
	models_order "backend/internal/models/order"
	"time"
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

func (object *orderServiceImplementation) cancelLimit(botModel *models_bot.BotModel) {
	botModel.OrderID = ""
	botModel.Deal = models_bot.BotDealModel{}

	object.balanceService().Release(botModel.ID)

	go func() {
		object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventBot,
			Data:  botModel,
		}
	}()
}

func (object *orderServiceImplementation) statusOpenLimit(orderModel *models_order.OrderModel, botModel *models_bot.BotModel) {
	botModel.Deal.PriceIn = orderModel.OriginalPrice
	botModel.Deal.AmountIn = orderModel.OriginalQuantity * orderModel.OriginalPrice
	botModel.Deal.AmountOut = orderModel.OriginalQuantity
	botModel.Deal.Status = enums_bot.DealStatusOpenLimit

	go func() {
		object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventBot,
			Data:  botModel,
		}
	}()
}

func (object *orderServiceImplementation) openLimit(orderModel *models_order.OrderModel, botModel *models_bot.BotModel) {
	priceOutFactor := (100 + botModel.Multiplier.Value*botModel.CurrentParam.PercentOut) / 100
	priceStopFactor := 0.0

	if botModel.CurrentParam.StopPercent > 0 {
		priceStopFactor = (100 - botModel.Multiplier.Value*botModel.CurrentParam.StopPercent) / 100
	}

	botModel.Deal.TimeIn = time.Now().UnixMilli()
	botModel.Deal.PriceIn = orderModel.AveragePrice
	botModel.Deal.AmountIn = orderModel.FilledQuantity * orderModel.AveragePrice
	botModel.Deal.AmountOut = orderModel.FilledQuantity
	botModel.Deal.PreparationPriceOut = object.getPriceOut(botModel.Deal.PriceIn, priceOutFactor, botModel.TickSizeFactor)
	botModel.Deal.PreparationPriceStop = object.getPriceStop(botModel.Deal.PriceIn, priceStopFactor, botModel.TickSizeFactor)

	if botModel.CurrentParam.StopTime > 0 {
		botModel.Deal.PreparationTimeOut = botModel.Deal.TimeIn + botModel.CurrentParam.StopTime
	}

	botModel.OrderID = ""
	botModel.Deal.Status = enums_bot.DealStatusOpen

	go func() {
		object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventBot,
			Data:  botModel,
		}

		if err := object.exchangeService().AddOutLimit(botModel, botModel.Deal.PreparationPriceOut, orderModel.FilledQuantity); err != nil {
			object.loggerService().Error().Printf("failed to add out limit: %v", err)
		}
	}()
}

func (object *orderServiceImplementation) closeLimit(orderModel *models_order.OrderModel, botModel *models_bot.BotModel) {
	botModel.Deal.TimeOut = time.Now().UnixMilli()
	botModel.Deal.PriceOut = orderModel.AveragePrice
	botModel.Deal.AmountOut = orderModel.FilledQuantity * orderModel.AveragePrice

	var delta float64

	if botModel.TradeDirection == enums.TradeDirectionLong {
		delta = botModel.Deal.AmountOut - botModel.Deal.AmountIn
	} else if botModel.TradeDirection == enums.TradeDirectionShort {
		delta = botModel.Deal.AmountIn - botModel.Deal.AmountOut
	}

	botModel.Deposit += delta

	if botModel.Deposit < 10 {
		botModel.Status = enums_bot.StatusStop
		botModel.Error = "not enough deposit"

		object.botRepositoryService().Remove(botModel.Symbol, botModel.ID)

		if err := object.botService().Update(botModel); err != nil {
			object.loggerService().Error().Printf("failed to update bot: %v", err)
		}
	}

	object.balanceService().Release(botModel.ID)

	botModelCopy := *botModel
	object.botService().GetAddDealChannel() <- &botModelCopy

	botModel.OrderID = ""
	botModel.Deal = models_bot.BotDealModel{}

	// if botModel.NextParam.PercentIn > 0 {
	// 	botModel.PrevParam = botModel.CurrentParam
	// 	botModel.CurrentParam = botModel.NextParam
	// 	botModel.NextParam = models_bot.BotParamModel{}
	// }

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBot,
		Data:  &botModelCopy,
	}
}

func (object *orderServiceImplementation) GetOrderChannel() chan *models_order.OrderModel {
	return object.orderChannel
}
