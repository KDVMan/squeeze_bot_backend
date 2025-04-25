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
				priceInFactor := (100 - botModel.Multiplier.Value*botModel.CurrentParam.PercentIn) / 100
				priceIn := object.getPriceIn(prevQuote, botModel.CurrentParam.Bind, priceInFactor, botModel.TickSizeFactor)

				if botModel.Deal.PreparationPriceIn != priceIn {
					if object.userService().GetAvailableBalance() < botModel.Deposit {
						log.Println("not enough balance", botModel.Symbol, "deposit", botModel.Deposit)
						continue
					}

					oldStatus := botModel.Deal.Status
					botModel.Deal.PreparationPriceIn = priceIn
					botModel.Deal.Status = enums_bot.DealStatusSendOpenLimit
					amount := botModel.Deposit / botModel.Deal.PreparationPriceIn

					go func() {
						object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
							Event: enums.WebsocketEventBot,
							Data:  botModel,
						}

						if oldStatus == enums_bot.DealStatusOpenLimit {
							if err := object.exchangeService().UpdateLimit(botModel, botModel.Deal.PreparationPriceIn, amount); err != nil {
								object.loggerService().Error().Printf("failed to update limit: %v", err)
							}
						} else {
							if err := object.exchangeService().AddInLimit(botModel, botModel.Deal.PreparationPriceIn, amount); err != nil {
								object.loggerService().Error().Printf("failed to add in limit: %v", err)

								object.botRepositoryService().Remove(botModel.Symbol, botModel.ID)

								botModel.Status = enums_bot.StatusStop
								botModel.Error = err.Error()

								if err = object.Update(botModel); err != nil {
									object.loggerService().Error().Printf("failed to update bot: %v", err)
								}
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

func (object *botServiceImplementation) GetDealChannel() chan string {
	return object.dealChannel
}
