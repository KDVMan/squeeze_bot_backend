package services_exchange

import (
	"backend/internal/enums"
	models_bot "backend/internal/models/bot"
	services_exchange_limit "backend/internal/services/exchange_limit"
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *exchangeServiceImplementation) AddOutMarket(botModel *models_bot.BotModel, amount float64) error {
	direction := futures.SideTypeSell
	positionSide := futures.PositionSideTypeLong

	if botModel.TradeDirection == enums.TradeDirectionShort {
		direction = futures.SideTypeBuy
		positionSide = futures.PositionSideTypeShort
	}

	botModel.OrderID = object.getOrderID(botModel.ID)

	_, err := object.client.NewCreateOrderService().
		Symbol(botModel.Symbol).
		Side(direction).
		PositionSide(positionSide).
		Type(futures.OrderTypeMarket).
		Quantity(fmt.Sprintf("%.*f", botModel.AmountPrecision, amount)).
		NewClientOrderID(botModel.OrderID).
		Do(context.Background())
	if err != nil {
		return err
	}

	if err = object.exchangeLimitService().Update(services_exchange_limit.GetLimits()); err != nil {
		return err
	}

	return nil
}
