package services_exchange

import (
	"backend/internal/enums"
	models_bot "backend/internal/models/bot"
	services_exchange_limit "backend/internal/services/exchange_limit"
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
	"log"
)

func (object *exchangeServiceImplementation) Limit(botModel *models_bot.BotModel, price float64, amount float64) error {
	direction := futures.SideTypeBuy
	positionSide := futures.PositionSideTypeLong

	if botModel.TradeDirection == enums.TradeDirectionShort {
		direction = futures.SideTypeSell
		positionSide = futures.PositionSideTypeShort
	}

	result, err := object.client.NewCreateOrderService().
		Symbol(botModel.Symbol).
		Side(direction).
		PositionSide(positionSide).
		Type(futures.OrderTypeLimit).
		TimeInForce(futures.TimeInForceTypeGTC).
		Quantity(fmt.Sprintf("%.*f", botModel.AmountPrecision, amount)).
		Price(fmt.Sprintf("%.*f", botModel.PricePrecision, price)).
		Do(context.Background())
	if err != nil {
		return err
	}

	log.Println("RESULT", result)
	object.dumpService().Dump(result)

	if err = object.exchangeLimitService().Update(services_exchange_limit.GetLimits()); err != nil {
		return err
	}

	// for _, balance := range result {
	// 	if balance.Asset == "USDT" {
	// 		amountBalance, err := strconv.ParseFloat(balance.Balance, 64)
	// 		if err != nil {
	// 			return 0, 0, err
	// 		}
	//
	// 		amountAvailableBalance, err := strconv.ParseFloat(balance.AvailableBalance, 64)
	// 		if err != nil {
	// 			return 0, 0, err
	// 		}
	//
	// 		return amountBalance, amountAvailableBalance, nil
	// 	}
	// }

	return nil
}
