package services_exchange

import (
	models_bot "backend/internal/models/bot"
	services_exchange_limit "backend/internal/services/exchange_limit"
	"context"
)

func (object *exchangeServiceImplementation) CancelLimit(botModel *models_bot.BotModel) error {
	_, err := object.client.NewCancelOrderService().
		Symbol(botModel.Symbol).
		OrigClientOrderID(botModel.OrderID).
		Do(context.Background())
	if err != nil {
		return err
	}

	if err = object.exchangeLimitService().Update(services_exchange_limit.GetLimits()); err != nil {
		return err
	}

	return nil
}
