package services_exchange

import (
	services_exchange_limit "backend/internal/services/exchange_limit"
	"context"
)

func (object *exchangeServiceImplementation) UserHedge() (bool, error) {
	result, err := object.client.NewGetPositionModeService().Do(context.Background())
	if err != nil {
		return false, err
	}

	if err = object.exchangeLimitService().Update(services_exchange_limit.GetLimits()); err != nil {
		return false, err
	}

	return result.DualSidePosition, nil
}
