package enums

import "github.com/go-playground/validator/v10"

type Interval string

const (
	Interval1m  Interval = "1m"
	Interval3m  Interval = "3m"
	Interval5m  Interval = "5m"
	Interval15m Interval = "15m"
	Interval30m Interval = "30m"
	Interval1h  Interval = "1h"
	Interval2h  Interval = "2h"
	Interval4h  Interval = "4h"
	Interval6h  Interval = "6h"
	Interval8h  Interval = "8h"
	Interval12h Interval = "12h"
	Interval1d  Interval = "1d"
	Interval3d  Interval = "3d"
	Interval1w  Interval = "1w"
	Interval1M  Interval = "1M"
)

func (enum Interval) String() string {
	return string(enum)
}

func IntervalSeconds(interval Interval) int64 {
	switch interval {
	case Interval1m:
		return 60
	case Interval3m:
		return 60 * 3
	case Interval5m:
		return 60 * 5
	case Interval15m:
		return 60 * 15
	case Interval30m:
		return 60 * 30
	case Interval1h:
		return 60 * 60
	case Interval2h:
		return 60 * 60 * 2
	case Interval4h:
		return 60 * 60 * 4
	case Interval6h:
		return 60 * 60 * 6
	case Interval8h:
		return 60 * 60 * 8
	case Interval12h:
		return 60 * 60 * 12
	case Interval1d:
		return 60 * 60 * 24
	case Interval3d:
		return 60 * 60 * 24 * 3
	case Interval1w:
		return 60 * 60 * 24 * 7
	case Interval1M:
		return 60 * 60 * 24 * 30
	default:
		panic("unknown interval: " + string(interval))
	}
}

func IntervalMilliseconds(interval Interval) int64 {
	return IntervalSeconds(interval) * 1000
}

func IntervalValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(Interval); ok {
		return enum.IntervalValid()
	}

	return false
}

func (enum Interval) IntervalValid() bool {
	switch enum {
	case Interval1m, Interval3m, Interval5m, Interval15m, Interval30m, Interval1h, Interval2h,
		Interval4h, Interval6h, Interval8h, Interval12h, Interval1d, Interval3d, Interval1w, Interval1M:
		return true
	default:
		return false
	}
}
