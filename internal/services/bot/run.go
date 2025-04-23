package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	models_quote "backend/internal/models/quote"
)

func (object *botServiceImplementation) RunChannel() {
	for botModel := range object.runChannel {
		// timeFrom, timeTo := models_quote.GetTimeRange(botModel.Interval, botModel.LimitQuotes)
		timeFrom, timeTo := models_quote.GetTimeRange(botModel.Interval, int64(object.configService().GetConfig().Binance.FuturesLimit))
		quoteRange := models_quote.GetRange(int64(object.configService().GetConfig().Binance.FuturesLimit), timeFrom, timeTo, enums.IntervalMilliseconds(enums.Interval1m))

		quotes, err := object.quoteService().LoadRange(botModel.Symbol, quoteRange)
		if err != nil {
			object.loggerService().Error().Printf("failed to load range: %v", err)
			continue
		}

		// что бы в репозиторий статус попал нормальный
		botModel.Status = enums_bot.StatusRun

		object.quoteRepositoryService().Add(botModel.Symbol, quotes)
		object.exchangeWebsocketService().SubscribeSymbol(botModel.Symbol)
		object.botRepositoryService().Add(botModel)

		if err = object.Status(&models_bot.StatusRequestModel{
			ID:     botModel.ID,
			Status: enums_bot.StatusRun,
		}); err != nil {
			object.loggerService().Error().Printf("failed to update bot status: %v", err)
			continue
		}
	}
}

func (object *botServiceImplementation) GetRunChannel() chan *models_bot.BotModel {
	return object.runChannel
}
