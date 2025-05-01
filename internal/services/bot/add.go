package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	enums_symbol "backend/internal/enums/symbol"
	models_bot "backend/internal/models/bot"
	models_channel "backend/internal/models/channel"
	services_helper "backend/pkg/services/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (object *botServiceImplementation) Add(request *models_bot.AddRequestModel) error {
	var botModel models_bot.BotModel
	var hash string

	if request.IsCalculator {
		hash = services_helper.MustConvertStringToMd5(
			fmt.Sprintf("hash | symbol:%s | window:%d | interval:%s | tradeDirection:%s",
				request.Symbol,
				request.Window,
				request.Interval.String(),
				request.TradeDirection.String(),
			),
		)
	} else {
		hash = services_helper.MustConvertStringToMd5(
			fmt.Sprintf("hash | symbol:%s | interval:%s | tradeDirection:%s | bind:%s | percentIn:%f | percentOut:%f | stopTime:%d | stopPercent:%f",
				request.Symbol,
				request.Interval.String(),
				request.TradeDirection.String(),
				request.Bind.String(),
				request.PercentIn,
				request.PercentOut,
				request.StopTime,
				request.StopPercent,
			),
		)
	}

	err := object.storageService().DB().Where("hash = ?", hash).First(&botModel).Error
	if err == nil {
		if request.IsCalculator {
			botModel.Deposit = request.Deposit
			botModel.IsReal = request.IsReal
			botModel.Window = request.Window
			botModel.LimitQuotes = request.LimitQuotes

			botModel.CurrentParam = models_bot.BotParamModel{
				Bind:         request.Bind,
				PercentIn:    request.PercentIn,
				PercentOut:   request.PercentOut,
				StopTime:     request.StopTime,
				StopPercent:  request.StopPercent,
				TriggerStart: request.TriggerStart,
			}

			if err = object.storageService().DB().Save(&botModel).Error; err != nil {
				return fmt.Errorf("failed to update calculator bot: %w", err)
			}

			if repositoryBotModel, ok := object.botRepositoryService().GetByID(botModel.ID, false); ok {
				if repositoryBotModel.Deal.Status == enums_bot.DealStatusOpen {
					repositoryBotModel.NextParam = botModel.CurrentParam
				} else {
					repositoryBotModel.CurrentParam = botModel.CurrentParam
				}

				object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
					Event: enums.WebsocketEventBot,
					Data:  repositoryBotModel,
				}
			} else {
				object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
					Event: enums.WebsocketEventBot,
					Data:  botModel,
				}
			}
		}

		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	symbolModel, err := object.symbolService().Load(request.Symbol, enums_symbol.SymbolStatusActive)
	if err != nil {
		return err
	}

	tickSizeFactor := 0
	tickSize := symbolModel.Limit.TickSize

	for tickSize < 1 {
		tickSize *= 10
		tickSizeFactor++
	}

	status := enums_bot.StatusAdd

	if request.IsCalculator {
		status = enums_bot.StatusRun
	}

	botModel = models_bot.BotModel{
		Hash:           hash,
		Deposit:        request.Deposit,
		IsReal:         request.IsReal,
		Symbol:         symbolModel.Symbol,
		Interval:       request.Interval,
		TradeDirection: request.TradeDirection,
		Window:         request.Window,
		LimitQuotes:    request.LimitQuotes,
		PrevParam:      models_bot.BotParamModel{},
		CurrentParam: models_bot.BotParamModel{
			Bind:         request.Bind,
			PercentIn:    request.PercentIn,
			PercentOut:   request.PercentOut,
			StopTime:     request.StopTime,
			StopPercent:  request.StopPercent,
			TriggerStart: request.TriggerStart,
		},
		NextParam:       models_bot.BotParamModel{},
		Multiplier:      models_bot.GetMultiplier(request.TradeDirection),
		TickSizeFactor:  tickSizeFactor,
		AmountPrecision: symbolModel.Limit.LeftPrecision,
		PricePrecision:  symbolModel.Limit.PricePrecision,
		Status:          status,
		IsCalculator:    request.IsCalculator,
	}

	if err = object.storageService().DB().Create(&botModel).Error; err != nil {
		return err
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	object.GetRunChannel() <- &botModel

	return nil
}
