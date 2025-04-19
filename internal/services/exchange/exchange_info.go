package services_exchange

import (
	services_exchange_limit "backend/internal/services/exchange_limit"
	"context"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *exchangeServiceImplementation) ExchangeInfo() ([]futures.Symbol, error) {
	result, err := object.client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	if err = object.exchangeLimitService().Create(result.RateLimits); err != nil {
		return nil, err
	}

	if err = object.exchangeLimitService().Update(services_exchange_limit.GetLimits()); err != nil {
		return nil, err
	}

	return result.Symbols, nil
}
