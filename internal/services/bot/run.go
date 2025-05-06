package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
	models_quote "backend/internal/models/quote"
)

func (object *botServiceImplementation) RunChannel() {
	for botModel := range object.runChannel {
		timeFrom, timeTo := models_quote.GetTimeRange(botModel.Interval, int64(object.configService().GetConfig().Binance.FuturesLimit))
		quoteRange := models_quote.GetRange(int64(object.configService().GetConfig().Binance.FuturesLimit), timeFrom, timeTo, enums.IntervalMilliseconds(enums.Interval1m))

		quotes, err := object.quoteService().LoadRange(botModel.Symbol, quoteRange)
		if err != nil {
			object.loggerService().Error().Printf("failed to load range: %v", err)
			continue
		}

		botModel.Status = enums_bot.StatusRun
		botModel.Error = ""
		botModel.Deal = models_bot.BotDealModel{}

		if err = object.storageService().DB().Save(&botModel).Error; err != nil {
			object.loggerService().Error().Printf("failed to save bot: %v", err)
			continue
		}

		object.quoteRepositoryService().Add(botModel.Symbol, quotes)
		object.exchangeWebsocketService().SubscribeSymbol(botModel.Symbol, false)
		object.botRepositoryService().Add(botModel)

		object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventBot,
			Data:  object.LoadByID(botModel.ID),
		}
	}
}

func (object *botServiceImplementation) GetRunChannel() chan *models_bot.BotModel {
	return object.runChannel
}
