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
	"time"
)

func (object *botServiceImplementation) AddCalculator(request *models_bot.AddCalculatorRequestModel) error {
	var botModel models_bot.BotModel

	hash := services_helper.MustConvertStringToMd5(
		fmt.Sprintf("hash | symbol:%s | window:%d | interval:%s | tradeDirection:%s",
			request.Symbol,
			request.Window,
			request.Interval.String(),
			request.TradeDirection.String(),
		),
	)

	err := object.storageService().DB().Where("hash = ?", hash).First(&botModel).Error
	if err == nil {
		botModel.Deposit = request.Deposit
		botModel.IsReal = true
		botModel.Window = request.Window
		botModel.LimitQuotes = request.LimitQuotes
		botModel.TimeUpdate = time.Now().UnixMilli()

		if request.PercentIn > 0 {
			botModel.CurrentParam = models_bot.BotParamModel{
				Bind:         request.Bind,
				PercentIn:    request.PercentIn,
				PercentOut:   request.PercentOut,
				StopTime:     request.StopTime,
				StopPercent:  request.StopPercent,
				TriggerStart: request.TriggerStart,
				MustUpdate:   true,
			}
		} else {
			botModel.CurrentParam = models_bot.BotParamModel{
				MustUpdate: true,
			}
		}

		if err = object.storageService().DB().Save(&botModel).Error; err != nil {
			return fmt.Errorf("failed to update calculator bot: %w", err)
		}

		if repositoryBotModel, ok := object.botRepositoryService().GetByID(botModel.ID, false); ok {
			repositoryBotModel.TimeUpdate = time.Now().UnixMilli()
			repositoryBotModel.NextParam = botModel.CurrentParam

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

	botModel = models_bot.BotModel{
		Hash:           hash,
		Deposit:        request.Deposit,
		IsReal:         true,
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
		IsCalculator:    true,
		TimeUpdate:      time.Now().UnixMilli(),
	}

	if err = object.storageService().DB().Create(&botModel).Error; err != nil {
		return err
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	if status == enums_bot.StatusAdd {
		object.GetRunChannel() <- &botModel
	}

	return nil
}
