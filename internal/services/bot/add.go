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

func (object *botServiceImplementation) Add(request *models_bot.AddRequestModel) error {
	var botModel models_bot.BotModel

	request.StopTime *= 60 * 1000

	hash := services_helper.MustConvertStringToMd5(
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

	err := object.storageService().DB().Where("hash = ?", hash).First(&botModel).Error
	if err == nil {
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
		Status:          enums_bot.StatusAdd,
		IsCalculator:    false,
		TimeUpdate:      time.Now().UnixMilli(),
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
