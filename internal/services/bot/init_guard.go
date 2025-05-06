package services_bot

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
)

func (object *botServiceImplementation) InitGuard() {
	symbols := []string{"BTCUSDT", "ETHUSDT"}

	for _, symbol := range symbols {
		timeFrom, timeTo := models_quote.GetTimeRange(enums.Interval1m, int64(object.configService().GetConfig().Binance.FuturesLimit))
		quoteRange := models_quote.GetRange(int64(object.configService().GetConfig().Binance.FuturesLimit), timeFrom, timeTo, enums.IntervalMilliseconds(enums.Interval1m))

		quotes, err := object.quoteService().LoadRange(symbol, quoteRange)
		if err != nil {
			object.loggerService().Error().Printf("failed to load range: %v", err)
		}

		object.quoteRepositoryService().Add(symbol, quotes)
		object.exchangeWebsocketService().SubscribeSymbol(symbol, true)
	}

	object.loggerService().Info().Printf("guard loaded")
}
