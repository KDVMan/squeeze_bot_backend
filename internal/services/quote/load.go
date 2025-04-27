package services_quote

import (
	enums_quote "backend/internal/enums/quote"
	models_bot "backend/internal/models/bot"
	models_deal "backend/internal/models/deal"
	models_quote "backend/internal/models/quote"
	services_helper "backend/pkg/services/helper"
	"fmt"
)

func (object *quoteServiceImplementation) Load(request *models_quote.QuoteRequestModel) (*models_quote.QuoteResponseModel, error) {
	var err error
	var quotes []*models_quote.QuoteModel
	hash := services_helper.MustConvertStringToMd5(fmt.Sprintf("hash | symbol:%s | interval:%s", request.Symbol, request.Interval.String()))
	response := &models_quote.QuoteResponseModel{}

	if request.Type == enums_quote.QuoteTypeInit {
		object.exchangeWebsocketService().SubscribeCurrentPrice(request.Symbol, request.Interval)

		if request.BotID > 0 {
			var deals []*models_deal.DealModel

			if err = object.storageService().DB().Where("bot_id = ?", request.BotID).Find(&deals).Error; err != nil {
				return nil, err
			}

			response.Deals = deals

			botModel, exists := object.botRepositoryService().GetByID(request.BotID, false)
			if exists {
				response.Bot = botModel
			} else {
				var botModelDB models_bot.BotModel

				if err = object.storageService().DB().First(&botModelDB, request.BotID).Error; err != nil {
					return nil, err
				}

				response.Bot = &botModelDB
			}
		}
	}

	if request.TimeEnd > 0 {
		quotes, err = object.loadLocal(hash, request)
		if err != nil {
			return nil, err
		}

		if len(quotes) < request.QuotesLimit {
			quotes, err = object.loadRemote(hash, request)
			if err != nil {
				return nil, err
			}
		}
	} else {
		quotes, err = object.loadRemote(hash, request)
		if err != nil {
			return nil, err
		}
	}

	response.TimeTo = request.TimeEnd
	response.Quotes = quotes

	return response, nil
}
