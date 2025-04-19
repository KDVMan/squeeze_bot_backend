package services_exchange

import (
	services_exchange_limit "backend/internal/services/exchange_limit"
	"context"
	"strconv"
)

func (object *exchangeServiceImplementation) UserBalance() (float64, float64, error) {
	result, err := object.client.NewGetBalanceService().Do(context.Background())
	if err != nil {
		return 0, 0, err
	}

	if err = object.exchangeLimitService().Update(services_exchange_limit.GetLimits()); err != nil {
		return 0, 0, err
	}

	for _, balance := range result {
		if balance.Asset == "USDT" {
			amountBalance, err := strconv.ParseFloat(balance.Balance, 64)
			if err != nil {
				return 0, 0, err
			}

			amountAvailableBalance, err := strconv.ParseFloat(balance.AvailableBalance, 64)
			if err != nil {
				return 0, 0, err
			}

			return amountBalance, amountAvailableBalance, nil
		}
	}

	return 0, 0, nil
}
