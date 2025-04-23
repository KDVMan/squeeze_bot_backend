package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
	"log"
)

func (object *botServiceImplementation) RunDealChannel() {
	for symbol := range object.dealChannel {
		botsModels, exists := object.botRepositoryService().GetBySymbol(symbol)
		if !exists {
			continue
		}

		quotes := object.quoteRepositoryService().GetBySymbol(symbol)
		currentQuote := quotes[len(quotes)-1]
		prevQuote := quotes[len(quotes)-2]

		for _, botModel := range botsModels {
			if botModel.Deal.StatusIsNull() {
				log.Println("NULL")

				priceInFactor := (100 - botModel.Multiplier.Value*botModel.CurrentParam.PercentIn) / 100
				priceIn := object.getPriceIn(prevQuote, botModel.CurrentParam.Bind, priceInFactor, botModel.TickSizeFactor)

				// priceOutFactor := (100 + botModel.Multiplier.Value*botModel.CurrentParam.PercentOut) / 100
				// priceStopFactor := 0.0

				// if botModel.CurrentParam.StopPercent > 0 {
				// 	priceStopFactor = (100 - botModel.Multiplier.Value*botModel.CurrentParam.StopPercent) / 100
				// }

				// botModel.Deal.TimeIn = time.Now().UnixMilli()
				// botModel.Deal.PriceIn = currentQuote.Price
				// botModel.Deal.AmountIn = botModel.Deposit / botModel.Deal.PriceIn
				// botModel.Deal.CalculatePriceOut = object.getPriceOut(botModel.Deal.PriceIn, priceOutFactor, botModel.TickSizeFactor)
				// botModel.Deal.CalculatePriceStop = object.getPriceStop(botModel.Deal.PriceIn, priceStopFactor, botModel.TickSizeFactor)

				// if botModel.CurrentParam.StopTime > 0 {
				// 	botModel.Deal.CalculateTimeOut = botModel.Deal.TimeIn + botModel.CurrentParam.StopTime*60*1000
				// }

				// botModel.Deal.Status = enums_bot.DealStatusSendOpenLimit

				// go func() {
				// 	if err := object.exchangeService().Limit(botModel); err != nil {
				// 		object.loggerService().Error().Printf("limit error: %v", err)
				// 	}
				// }()

				// object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
				// 	Event: enums.WebsocketEventBot,
				// 	Data:  botModel,
				// }

				// log.Println("in", botModel.Status)

				if botModel.Deal.PreparationPriceIn != priceIn {
					botModel.Deal.PreparationPriceIn = priceIn
					botModel.Deal.Status = enums_bot.DealStatusSendOpenLimit
					amount := botModel.Deposit / botModel.Deal.PreparationPriceIn

					log.Println("PreparationPriceIn", botModel.Deal.PreparationPriceIn, "amount", amount, "status", botModel.Deal.Status)

					go func() {
						object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
							Event: enums.WebsocketEventBot,
							Data:  botModel,
						}

						if err := object.exchangeService().Limit(botModel, botModel.Deal.PreparationPriceIn, amount); err != nil {
							object.loggerService().Error().Printf("failed to open limit: %v", err)
						}
					}()
				}
			} else if botModel.Deal.Status == enums_bot.DealStatusOpenLimit {
			} else if botModel.Deal.Status == enums_bot.DealStatusOpen && object.checkCloseDeal(botModel, currentQuote.Price) {
				botModelCopy := *botModel
				object.GetAddDealChannel() <- &botModelCopy

				botModel.Deal = models_bot.BotDealModel{}

				object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
					Event: enums.WebsocketEventBot,
					Data:  &botModelCopy,
				}

				log.Println("out", botModel.Status)
			}
		}
	}
}

// func (object *botServiceImplementation) checkOpenDeal(botModel *models_bot.BotModel, currentQuote *models_quote.QuoteModel) bool {
// 	if botModel.TradeDirection == enums.TradeDirectionLong {
// 		return currentQuote.Price <= botModel.PreparationDeal.PriceIn
// 	} else if botModel.TradeDirection == enums.TradeDirectionShort {
// 		return currentQuote.Price >= botModel.PreparationDeal.PriceIn
// 	}
//
// 	return false
// }

func (object *botServiceImplementation) checkCloseDeal(botModel *models_bot.BotModel, currentPrice float64) bool {
	// now := time.Now().UnixMilli()
	//
	// switch botModel.TradeDirection {
	// case enums.TradeDirectionLong:
	// 	if currentPrice >= botModel.Deal.CalculatePriceOut {
	// 		botModel.Deal.PriceOut = currentPrice
	// 	} else if botModel.Deal.CalculatePriceStop > 0 && currentPrice <= botModel.Deal.CalculatePriceStop {
	// 		botModel.Deal.PriceOut = currentPrice
	// 		botModel.Deal.IsStopPercent = true
	// 	} else if botModel.Deal.CalculateTimeOut > 0 && now >= botModel.Deal.CalculateTimeOut {
	// 		botModel.Deal.PriceOut = currentPrice
	// 		botModel.Deal.IsStopTime = true
	// 	} else {
	// 		return false
	// 	}
	// case enums.TradeDirectionShort:
	// 	if currentPrice <= botModel.Deal.CalculatePriceOut {
	// 		botModel.Deal.PriceOut = currentPrice
	// 	} else if botModel.Deal.CalculatePriceStop > 0 && currentPrice >= botModel.Deal.CalculatePriceStop {
	// 		botModel.Deal.PriceOut = currentPrice
	// 		botModel.Deal.IsStopPercent = true
	// 	} else if botModel.Deal.CalculateTimeOut > 0 && now >= botModel.Deal.CalculateTimeOut {
	// 		botModel.Deal.PriceOut = currentPrice
	// 		botModel.Deal.IsStopTime = true
	// 	} else {
	// 		return false
	// 	}
	// default:
	// 	return false
	// }
	//
	// botModel.Deal.TimeOut = now
	// botModel.Deal.AmountOut = botModel.Deal.AmountIn * botModel.Deal.PriceOut
	// botModel.Deal.Status = enums_bot.DealStatusClose
	//
	// if botModel.NextParam.PercentIn > 0 {
	// 	botModel.PrevParam = botModel.CurrentParam
	// 	botModel.CurrentParam = botModel.NextParam
	// 	botModel.NextParam = models_bot.BotParamModel{}
	// }

	return true
}

func (object *botServiceImplementation) GetDealChannel() chan string {
	return object.dealChannel
}
