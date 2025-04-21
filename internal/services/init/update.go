package services_init

import (
	"backend/internal/enums"
	enums_symbol "backend/internal/enums/symbol"
	models_channel "backend/internal/models/channel"
	models_init "backend/internal/models/init"
	models_symbol "backend/internal/models/symbol"
)

func (object *initServiceImplementation) Update(request *models_init.UpdateRequestModel) (*models_init.InitModel, error) {
	initModel, err := object.Load()
	if err != nil {
		return nil, err
	}

	initModel.Symbol = request.Symbol
	initModel.Intervals = request.Intervals
	initModel.BotID = request.BotID
	initModel.BotSortColumn = request.BotSortColumn
	initModel.BotSortDirection = request.BotSortDirection

	symbolModel, err := object.symbolService().Load(request.Symbol, enums_symbol.SymbolStatusActive)
	if err != nil {
		return nil, err
	}

	initModel.Precision = symbolModel.Limit.Precision
	initModel.LeverageLevel = symbolModel.Leverage.Level
	initModel.LeverageType = symbolModel.Leverage.Type

	result := object.storageService().DB().Save(&initModel)
	if result.Error != nil {
		return nil, err
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventLeverage,
		Data: &models_symbol.SymbolLeverageModel{
			Level: symbolModel.Leverage.Level,
			Type:  symbolModel.Leverage.Type,
		},
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventBotList,
		Data:  object.botService().LoadAll(),
	}

	return initModel, nil
}
