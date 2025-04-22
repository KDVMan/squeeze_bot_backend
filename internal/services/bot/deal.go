package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	services_helper "backend/pkg/services/helper"
	"log"
	"time"
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

		// log.Printf(
		// 	"current, timeOpen: %s, timeClose: %s, price: %f,closed: %v\n",
		// 	services_helper.MustConvertUnixMillisecondsToString(currentQuote.TimeOpen),
		// 	services_helper.MustConvertUnixMillisecondsToString(currentQuote.TimeClose),
		// 	currentQuote.Price,
		// 	currentQuote.IsClosed,
		// )
		//
		// log.Printf(
		// 	"prev, timeOpen: %s, timeClose: %s, price: %f,closed: %v\n\n",
		// 	services_helper.MustConvertUnixMillisecondsToString(prevQuote.TimeOpen),
		// 	services_helper.MustConvertUnixMillisecondsToString(prevQuote.TimeClose),
		// 	prevQuote.Price,
		// 	prevQuote.IsClosed,
		// )

		for _, botModel := range botsModels {
			priceInFactor := (100 - botModel.Multiplier.Value*botModel.CurrentParam.PercentIn) / 100
			priceIn := object.getPriceIn(prevQuote, botModel.CurrentParam.Bind, priceInFactor, botModel.TickSizeFactor)

			if botModel.Deal.Status == enums_bot.DealStatusOpen {
				now := time.Now().UnixMilli()

				if botModel.TradeDirection == enums.TradeDirectionLong {
					if currentQuote.Price >= botModel.Deal.CalculatePriceOut {
						botModel.Deal.TimeOut = now
						botModel.Deal.PriceOut = currentQuote.Price
						botModel.Deal.Status = enums_bot.DealStatusClose
					} else if botModel.Deal.CalculatePriceStop > 0 && currentQuote.Price <= botModel.Deal.CalculatePriceStop {
						botModel.Deal.TimeOut = now
						botModel.Deal.PriceOut = currentQuote.Price
						botModel.Deal.IsStopPercent = true
						botModel.Deal.Status = enums_bot.DealStatusClose
					} else if botModel.Deal.CalculateTimeOut > 0 && now >= botModel.Deal.CalculateTimeOut {
						botModel.Deal.TimeOut = now
						botModel.Deal.PriceOut = currentQuote.Price
						botModel.Deal.IsStopTime = true
						botModel.Deal.Status = enums_bot.DealStatusClose
					}
				} else if botModel.TradeDirection == enums.TradeDirectionShort {
					if currentQuote.Price <= botModel.Deal.CalculatePriceOut {
						botModel.Deal.TimeOut = now
						botModel.Deal.PriceOut = currentQuote.Price
						botModel.Deal.Status = enums_bot.DealStatusClose
					} else if botModel.Deal.CalculatePriceStop > 0 && currentQuote.Price >= botModel.Deal.CalculatePriceStop {
						botModel.Deal.TimeOut = now
						botModel.Deal.PriceOut = currentQuote.Price
						botModel.Deal.IsStopPercent = true
						botModel.Deal.Status = enums_bot.DealStatusClose
					} else if botModel.Deal.CalculateTimeOut > 0 && now >= botModel.Deal.CalculateTimeOut {
						botModel.Deal.TimeOut = now
						botModel.Deal.PriceOut = currentQuote.Price
						botModel.Deal.IsStopTime = true
						botModel.Deal.Status = enums_bot.DealStatusClose
					}
				}

				if botModel.Deal.Status == enums_bot.DealStatusClose {
					botModelCopy := *botModel
					object.GetAddDealChannel() <- &botModelCopy

					botModel.Deal = models_bot.BotDealModel{}
				}
			} else if botModel.Deal.IsNull() && object.shouldEnterDeal(botModel.TradeDirection, currentQuote.Price, priceIn) {
				priceOutFactor := (100 + botModel.Multiplier.Value*botModel.CurrentParam.PercentOut) / 100
				priceStopFactor := 0.0

				if botModel.CurrentParam.StopPercent > 0 {
					priceStopFactor = (100 - botModel.Multiplier.Value*botModel.CurrentParam.StopPercent) / 100
				}

				botModel.Deal.TimeIn = time.Now().UnixMilli()
				botModel.Deal.PriceIn = currentQuote.Price
				botModel.Deal.CalculatePriceOut = object.getPriceOut(botModel.Deal.PriceIn, priceOutFactor, botModel.TickSizeFactor)
				botModel.Deal.CalculatePriceStop = object.getPriceStop(botModel.Deal.PriceIn, priceStopFactor, botModel.TickSizeFactor)

				if botModel.CurrentParam.StopTime > 0 {
					botModel.Deal.CalculateTimeOut = botModel.Deal.TimeIn + botModel.CurrentParam.StopTime*60*1000
				}

				botModel.Deal.Status = enums_bot.DealStatusOpen

				log.Println("IN, stop: ", services_helper.MustConvertUnixMillisecondsToString(botModel.Deal.CalculateTimeOut))
				object.dumpService().Dump(botModel.Deal)
			}
		}
	}
}

func (object *botServiceImplementation) shouldEnterDeal(tradeDirection enums.TradeDirection, currentPrice, priceIn float64) bool {
	if tradeDirection == enums.TradeDirectionLong {
		return currentPrice <= priceIn
	} else if tradeDirection == enums.TradeDirectionShort {
		return currentPrice >= priceIn
	}

	return false
}

func (object *botServiceImplementation) GetDealChannel() chan string {
	return object.dealChannel
}
