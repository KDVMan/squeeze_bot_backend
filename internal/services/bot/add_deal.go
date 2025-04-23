package services_bot

import (
	"backend/internal/enums"
	models_bot "backend/internal/models/bot"
	models_deal "backend/internal/models/deal"
)

func (object *botServiceImplementation) RunAddDealChannel() {
	for botModel := range object.addDealChannel {
		dealModel := &models_deal.DealModel{
			BotID:         botModel.ID,
			IsReal:        botModel.IsReal,
			TimeIn:        botModel.Deal.TimeIn,
			TimeOut:       botModel.Deal.TimeOut,
			PriceIn:       botModel.Deal.PriceIn,
			AmountIn:      botModel.Deal.AmountIn,
			PriceOut:      botModel.Deal.PriceOut,
			AmountOut:     botModel.Deal.AmountOut,
			IsStopTime:    botModel.Deal.IsStopTime,
			IsStopPercent: botModel.Deal.IsStopPercent,
			ProfitPercent: object.calculateProfit(botModel),
			Param: models_deal.DealParamModel{
				CalculatorId: botModel.CurrentParam.CalculatorId,
				Bind:         botModel.CurrentParam.Bind,
				PercentIn:    botModel.CurrentParam.PercentIn,
				PercentOut:   botModel.CurrentParam.PercentOut,
				StopTime:     botModel.CurrentParam.StopTime,
				StopPercent:  botModel.CurrentParam.StopPercent,
			},
		}

		if err := object.storageService().DB().Create(dealModel).Error; err != nil {
			object.loggerService().Error().Printf("failed to save deal: %v", err)
			continue
		}
	}
}

func (object *botServiceImplementation) calculateProfit(botModel *models_bot.BotModel) float64 {
	if botModel.TradeDirection == enums.TradeDirectionShort {
		return 100/botModel.Deal.PriceOut*botModel.Deal.PriceIn - 100 - (100+100/botModel.Deal.PriceOut*botModel.Deal.PriceIn)*object.commission/100
	} else if botModel.TradeDirection == enums.TradeDirectionLong {
		return 100/botModel.Deal.PriceIn*botModel.Deal.PriceOut - 100 - (100+100/botModel.Deal.PriceIn*botModel.Deal.PriceOut)*object.commission/100
	}

	return 0
}

func (object *botServiceImplementation) GetAddDealChannel() chan *models_bot.BotModel {
	return object.addDealChannel
}
