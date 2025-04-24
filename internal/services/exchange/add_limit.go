package services_exchange

import (
	"backend/internal/enums"
	models_bot "backend/internal/models/bot"
	services_exchange_limit "backend/internal/services/exchange_limit"
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *exchangeServiceImplementation) AddLimit(botModel *models_bot.BotModel, price float64, amount float64) error {
	direction := futures.SideTypeBuy
	positionSide := futures.PositionSideTypeLong

	if botModel.TradeDirection == enums.TradeDirectionShort {
		direction = futures.SideTypeSell
		positionSide = futures.PositionSideTypeShort
	}

	orderID := object.getOrderID(botModel.ID)

	_, err := object.client.NewCreateOrderService().
		Symbol(botModel.Symbol).
		Side(direction).
		PositionSide(positionSide).
		Type(futures.OrderTypeLimit).
		TimeInForce(futures.TimeInForceTypeGTC).
		Quantity(fmt.Sprintf("%.*f", botModel.AmountPrecision, amount)).
		Price(fmt.Sprintf("%.*f", botModel.PricePrecision, price)).
		NewClientOrderID(orderID).
		Do(context.Background())
	if err != nil {
		return err
	}

	if err = object.exchangeLimitService().Update(services_exchange_limit.GetLimits()); err != nil {
		return err
	}

	return nil
}
