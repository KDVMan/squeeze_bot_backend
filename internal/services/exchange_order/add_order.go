package services_exchange_order

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *exchangeOrderServiceImplementation) AddOrder(botModel *models_bot.BotModel, price float64, amount float64) error {
	direction := futures.SideTypeBuy
	positionSide := futures.PositionSideTypeLong

	if botModel.TradeDirection == enums.TradeDirectionShort {
		direction = futures.SideTypeSell
		positionSide = futures.PositionSideTypeShort
	}

	orderID := object.getOrderID(botModel.ID)

	request := futures.NewOrderPlaceWsRequest().
		Symbol(botModel.Symbol).
		Side(direction).
		PositionSide(positionSide).
		Type(futures.OrderTypeLimit).
		TimeInForce(futures.TimeInForceTypeGTC).
		Quantity(fmt.Sprintf("%.*f", botModel.AmountPrecision, amount)).
		Price(fmt.Sprintf("%.*f", botModel.PricePrecision, price)).
		NewClientOrderID(orderID)

	if err := object.orderPlaceService.Do(orderID, request); err != nil {
		object.mutex.Lock()
		botModel.OrderID = ""
		object.mutex.Unlock()

		return err
	}

	object.mutex.Lock()
	botModel.OrderID = orderID
	botModel.Deal.Status = enums_bot.DealStatusSendOpenLimitWs
	object.mutex.Unlock()

	object.dumpService().Dump(botModel)

	return nil
}
