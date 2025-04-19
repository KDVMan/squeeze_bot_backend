package services_exchange_limit

import (
	enums_exchange_limit "backend/internal/enums/exchange_limit"
	models_exchange_limit "backend/internal/models/exchange_limit"
	"github.com/adshao/go-binance/v2/futures"
)

func (object *exchangeLimitServiceImplementation) Create(limits []futures.RateLimit) error {
	for _, limit := range limits {
		exchangeLimitModel := models_exchange_limit.ExchangeLimitModel{
			Type:           convertType(limit.RateLimitType),
			Interval:       convertInterval(limit.Interval),
			IntervalNumber: limit.IntervalNum,
			Total:          limit.Limit,
			TotalLeft:      limit.Limit,
		}

		err := object.storageService().DB().
			Where("type = ? AND interval = ? AND interval_number = ?", exchangeLimitModel.Type, exchangeLimitModel.Interval, exchangeLimitModel.IntervalNumber).
			Assign(exchangeLimitModel).
			FirstOrCreate(&exchangeLimitModel).
			Error

		if err != nil {
			return err
		}
	}

	return nil
}

func convertType(input string) enums_exchange_limit.RateType {
	switch input {
	case "REQUEST_WEIGHT":
		return enums_exchange_limit.RateTypeWeight
	case "ORDERS":
		return enums_exchange_limit.RateTypeOrder
	default:
		return enums_exchange_limit.RateTypeUnknown
	}
}

func convertInterval(input string) enums_exchange_limit.RateInterval {
	switch input {
	case "SECOND":
		return enums_exchange_limit.RateIntervalSecond
	case "MINUTE":
		return enums_exchange_limit.RateIntervalMinute
	case "HOUR":
		return enums_exchange_limit.RateIntervalHour
	case "DAY":
		return enums_exchange_limit.RateIntervalDay
	default:
		return enums_exchange_limit.RateIntervalUnknown
	}
}
