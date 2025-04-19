package services_exchange_limit

import (
	models_exchange_limit "backend/internal/models/exchange_limit"
)

func (object *exchangeLimitServiceImplementation) Load() ([]*models_exchange_limit.ExchangeLimitModel, error) {
	var exchangeLimitModel []*models_exchange_limit.ExchangeLimitModel

	if err := object.storageService().DB().Find(&exchangeLimitModel).Error; err != nil {
		return nil, err
	}

	return exchangeLimitModel, nil
}
