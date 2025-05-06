package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
	"log"
	"time"
)

func (object *botServiceImplementation) RunDealChannel() {
	for {
		select {
		case symbol := <-object.dealChannel:
			object.handleDeal(symbol)

		case event := <-object.botEventChannel:
			object.handleBotEvent(event)
		}
	}
}

func (object *botServiceImplementation) handleDeal(symbol string) {
	botsModels, exists := object.botRepositoryService().GetBySymbol(symbol)
	if !exists {
		return
	}

	quotes := object.quoteRepositoryService().GetBySymbol(symbol)
	currentQuote := quotes[len(quotes)-1]
	prevQuote := quotes[len(quotes)-2]

	for _, botModel := range botsModels {
		if botModel.Deal.StatusIsNull() {
			var percentToPriceIn float64

			object.UpdateParam(botModel)

			priceInFactor := (100 - botModel.Multiplier.Value*botModel.CurrentParam.PercentIn) / 100
			priceIn := object.getPriceIn(prevQuote, botModel.CurrentParam.Bind, priceInFactor, botModel.TickSizeFactor)

			if priceIn <= 0 {
				continue
			}

			if botModel.TradeDirection == enums.TradeDirectionLong {
				percentToPriceIn = ((currentQuote.Price - priceIn) / priceIn) * 100
			} else if botModel.TradeDirection == enums.TradeDirectionShort {
				percentToPriceIn = ((priceIn - currentQuote.Price) / priceIn) * 100
			}

			if percentToPriceIn <= botModel.CurrentParam.TriggerStart && botModel.Deal.Status != enums_bot.DealStatusOpenLimit {
				if !object.balanceService().Reserve(botModel.ID, botModel.Deposit) {
					continue
				}

				if object.isGuardActive(botModel.Symbol, botModel.TradeDirection) {
					object.loggerService().Info().Printf("[GUARD] %s: skipping deal due to active stop %s",
						botModel.Symbol,
						botModel.TradeDirection,
					)

					continue
				}

				botModel.Deal.PreparationPriceIn = priceIn
				botModel.Deal.Status = enums_bot.DealStatusSendOpenLimit
				botModel.Deal.TriggerTime = time.Now().UnixMilli()
				amount := botModel.Deposit / botModel.Deal.PreparationPriceIn

				if amount <= 0 || botModel.Deposit <= 0 || botModel.Deal.PreparationPriceIn <= 0 {
					log.Printf("invalid amount: deposit=%.4f, price=%.8f -> amount=%.8f", botModel.Deposit, botModel.Deal.PreparationPriceIn, amount)
					object.balanceService().Release(botModel.ID)
					continue
				}

				object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
					Event: enums.WebsocketEventBot,
					Data:  botModel,
				}

				if err := object.exchangeService().AddInLimit(botModel, botModel.Deal.PreparationPriceIn, amount); err != nil {
					object.loggerService().Error().Printf("failed to add in limit: %v", err)

					object.balanceService().Release(botModel.ID)
					object.botRepositoryService().Remove(botModel.Symbol, botModel.ID)

					botModel.Status = enums_bot.StatusStop
					botModel.Error = err.Error()

					if err = object.Update(botModel); err != nil {
						object.loggerService().Error().Printf("failed to update bot: %v", err)
					}
				}
			} else if percentToPriceIn > botModel.CurrentParam.TriggerStart && botModel.Deal.Status == enums_bot.DealStatusOpenLimit {
				now := time.Now().UnixMilli()

				if now-botModel.Deal.TriggerTime < 2000 {
					continue
				}

				botModel.Deal.Status = enums_bot.DealStatusSendCancel

				go func() {
					object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
						Event: enums.WebsocketEventBot,
						Data:  botModel,
					}

					if err := object.exchangeService().CancelLimit(botModel); err != nil {
						object.loggerService().Error().Printf("failed to cancel limit : %v", err)

						object.balanceService().Release(botModel.ID)
						object.botRepositoryService().Remove(botModel.Symbol, botModel.ID)

						botModel.Status = enums_bot.StatusStop
						botModel.Error = err.Error()

						if err = object.Update(botModel); err != nil {
							object.loggerService().Error().Printf("failed to update bot: %v", err)
						}
					}
				}()
			}
		} else if botModel.Deal.Status == enums_bot.DealStatusOpen && object.checkStop(botModel, currentQuote.Price) {
			botModel.Deal.Status = enums_bot.DealStatusSendClose

			go func() {
				object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
					Event: enums.WebsocketEventBot,
					Data:  botModel,
				}

				if err := object.exchangeService().AddOutMarket(botModel, botModel.Deal.AmountOut); err != nil {
					object.loggerService().Error().Printf("failed to add out market: %v", err)
				}
			}()
		}
	}
}

func (object *botServiceImplementation) handleBotEvent(event *models_bot.BotEventModel) {
	botModel, exists := object.botRepositoryService().GetBySymbolAndID(event.Symbol, event.BotID)
	if !exists {
		return
	}

	switch event.Type {
	case enums_bot.BotEventStatusOpenLimit:
		botModel.Deal.PriceIn = event.Order.OriginalPrice
		botModel.Deal.AmountIn = event.Order.OriginalQuantity * event.Order.OriginalPrice
		botModel.Deal.AmountOut = event.Order.OriginalQuantity
		botModel.Deal.Status = enums_bot.DealStatusOpenLimit

	case enums_bot.BotEventCancelLimit:
		botModel.OrderID = ""
		botModel.Deal = models_bot.BotDealModel{}

		object.balanceService().Release(botModel.ID)

	case enums_bot.BotEventOpenLimit:
		priceOutFactor := (100 + botModel.Multiplier.Value*botModel.CurrentParam.PercentOut) / 100
		priceStopFactor := 0.0

		if botModel.CurrentParam.StopPercent > 0 {
			priceStopFactor = (100 - botModel.Multiplier.Value*botModel.CurrentParam.StopPercent) / 100
		}

		botModel.Deal.TimeIn = time.Now().UnixMilli()
		botModel.Deal.PriceIn = event.Order.AveragePrice
		botModel.Deal.AmountIn = event.Order.FilledQuantity * event.Order.AveragePrice
		botModel.Deal.AmountOut = event.Order.FilledQuantity
		botModel.Deal.PreparationPriceOut = object.getPriceOut(botModel.Deal.PriceIn, priceOutFactor, botModel.TickSizeFactor)
		botModel.Deal.PreparationPriceStop = object.getPriceStop(botModel.Deal.PriceIn, priceStopFactor, botModel.TickSizeFactor)

		if botModel.CurrentParam.StopTime > 0 {
			botModel.Deal.PreparationTimeOut = botModel.Deal.TimeIn + botModel.CurrentParam.StopTime
		}

		botModel.OrderID = ""
		botModel.Deal.Status = enums_bot.DealStatusOpen

		go func() {
			if err := object.exchangeService().AddOutLimit(botModel, botModel.Deal.PreparationPriceOut, event.Order.FilledQuantity); err != nil {
				object.loggerService().Error().Printf("failed to add out limit: %v", err)
			}
		}()

	case enums_bot.BotEventCloseLimit:
		botModel.Deal.TimeOut = time.Now().UnixMilli()
		botModel.Deal.PriceOut = event.Order.AveragePrice
		botModel.Deal.AmountOut = event.Order.FilledQuantity * event.Order.AveragePrice

		var delta float64

		if botModel.TradeDirection == enums.TradeDirectionLong {
			delta = botModel.Deal.AmountOut - botModel.Deal.AmountIn
		} else {
			delta = botModel.Deal.AmountIn - botModel.Deal.AmountOut
		}

		botModel.Deposit += delta

		object.balanceService().Release(botModel.ID)
		object.balanceService().UpdateBalance(delta)

		if botModel.Deposit < 10 {
			botModel.Status = enums_bot.StatusStop
			botModel.Error = "not enough deposit"

			object.botRepositoryService().Remove(botModel.Symbol, botModel.ID)

			if err := object.Update(botModel); err != nil {
				object.loggerService().Error().Printf("failed to update bot: %v", err)
			}
		}

		botCopy := *botModel
		object.GetAddDealChannel() <- &botCopy

		botModel.OrderID = ""
		botModel.Deal = models_bot.BotDealModel{}
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBot,
		Data:  botModel,
	}
}

func (object *botServiceImplementation) checkStop(botModel *models_bot.BotModel, currentPrice float64) bool {
	now := time.Now().UnixMilli()

	switch botModel.TradeDirection {
	case enums.TradeDirectionLong:
		if botModel.Deal.PreparationPriceStop > 0 && currentPrice <= botModel.Deal.PreparationPriceStop {
			botModel.Deal.IsStopPercent = true
		} else if botModel.Deal.PreparationTimeOut > 0 && now >= botModel.Deal.PreparationTimeOut {
			botModel.Deal.IsStopTime = true
		} else {
			return false
		}
	case enums.TradeDirectionShort:
		if botModel.Deal.PreparationPriceStop > 0 && currentPrice >= botModel.Deal.PreparationPriceStop {
			botModel.Deal.IsStopPercent = true
		} else if botModel.Deal.PreparationTimeOut > 0 && now >= botModel.Deal.PreparationTimeOut {
			botModel.Deal.IsStopTime = true
		} else {
			return false
		}
	default:
		return false
	}

	return true
}

func (object *botServiceImplementation) UpdateParam(botModel *models_bot.BotModel) {
	if botModel.Deal.Status != "" && botModel.Deal.Status != enums_bot.DealStatusNull {
		return
	}

	if botModel.NextParam.PercentIn > 0 || botModel.NextParam.MustUpdate {
		botModel.PrevParam = botModel.CurrentParam
		botModel.CurrentParam = botModel.NextParam
		botModel.NextParam = models_bot.BotParamModel{}

		object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventBot,
			Data:  botModel,
		}
	}
}

func (object *botServiceImplementation) GetDealChannel() chan string {
	return object.dealChannel
}

func (object *botServiceImplementation) GetBotEventChannel() chan *models_bot.BotEventModel {
	return object.botEventChannel
}
