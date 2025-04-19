package services_exchange

import (
	services_exchange_limit "backend/internal/services/exchange_limit"
	"context"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *exchangeServiceImplementation) Kline(symbol string, interval string, timeEnd int64, limit int) ([]*futures.Kline, error) {
	query := object.client.NewKlinesService().
		Symbol(symbol).
		Interval(interval).
		Limit(limit)

	if timeEnd > 0 {
		query = query.EndTime(timeEnd)
	}

	result, err := query.Do(context.Background())
	if err != nil {
		return nil, err
	}

	if err = object.exchangeLimitService().Update(services_exchange_limit.GetLimits()); err != nil {
		return nil, err
	}

	return result, nil
}
