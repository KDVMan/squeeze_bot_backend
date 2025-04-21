package services_bot

import (
	"backend/internal/enums"
	models_bot "backend/internal/models/bot"
	models_quote "backend/internal/models/quote"
	services_quote_builder "backend/internal/services/quote_builder"
	services_helper "backend/pkg/services/helper"
	"log"
)

func (object *botServiceImplementation) RunDealChannel() {
	for symbol := range object.dealChannel {
		botsModels, exists := object.botRepositoryService().GetBySymbol(symbol)
		if !exists {
			continue
		}

		quotes := object.quoteRepositoryService().GetBySymbol(symbol)

		for _, botModel := range botsModels {
			if botModel.CurrentParam.PercentIn <= 0 {
				continue
			}

			var direction models_bot.DirectionModel

			if botModel.TradeDirection == enums.TradeDirectionShort {
				direction = models_bot.DirectionModel{Multiplier: -1, MinKeyName: enums.BindHigh, MaxKeyName: enums.BindLow}
			} else {
				direction = models_bot.DirectionModel{Multiplier: 1, MinKeyName: enums.BindLow, MaxKeyName: enums.BindHigh}
			}

			priceInFactor := (100 - direction.Multiplier*botModel.CurrentParam.PercentIn) / 100
			// priceOutFactor := (100 + direction.Multiplier*botModel.CurrentParam.PercentOut) / 100
			tickSizeFactor := 0
			tickSize := botModel.TickSize

			for tickSize < 1 {
				tickSize *= 10
				tickSizeFactor++
			}

			quoteBuilderService := services_quote_builder.NewQuoteBuilderService(botModel.Interval, enums.Interval1m)
			quoteBuild := quoteBuilderService.Build(quotes[len(quotes)-1])
			priceIn := object.calculatePriceIn(quoteBuild, botModel.CurrentParam.Bind, priceInFactor, tickSizeFactor)

			if botModel.InDeal {

			} else {
				if object.calculatePriceBind(direction.MinKeyName, quoteBuild) < priceIn {
					log.Println("Buy", symbol, object.calculatePriceBind(direction.MinKeyName, quoteBuild), " < ", priceIn)
				}
			}
		}
	}
}

func (object *botServiceImplementation) calculatePriceIn(quote *models_quote.QuoteModel, bind enums.Bind, priceInFactor float64, tickSizeFactor int) float64 {
	priceBind := object.calculatePriceBind(bind, quote)
	priceIn := priceBind * priceInFactor

	return services_helper.Floor(priceIn, tickSizeFactor)
}

func (object *botServiceImplementation) calculatePriceBind(bind enums.Bind, quote *models_quote.QuoteModel) float64 {
	switch bind {
	case enums.BindLow:
		return quote.PriceLow
	case enums.BindHigh:
		return quote.PriceHigh
	case enums.BindOpen:
		return quote.PriceOpen
	case enums.BindClose:
		return quote.PriceClose
	case enums.BindMhl:
		return (quote.PriceHigh + quote.PriceLow) / 2
	case enums.BindMoc:
		return (quote.PriceOpen + quote.PriceClose) / 2
	default:
		return 0
	}
}

func (object *botServiceImplementation) GetDealChannel() chan string {
	return object.dealChannel
}
