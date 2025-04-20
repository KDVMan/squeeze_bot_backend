package enums_bot

import "github.com/go-playground/validator/v10"

type SortColumn string

const (
	SortColumnId             SortColumn = "id"
	SortColumnSymbol         SortColumn = "symbol"
	SortColumnInterval       SortColumn = "interval"
	SortColumnTradeDirection SortColumn = "tradeDirection"
	SortColumnWindow         SortColumn = "window"
	SortColumnStatus         SortColumn = "status"
	SortColumnBind           SortColumn = "bind"
	SortColumnPercentIn      SortColumn = "percentIn"
	SortColumnPercentOut     SortColumn = "percentOut"
	SortColumnStopTime       SortColumn = "stopTime"
	SortColumnStopPercent    SortColumn = "stopPercent"
	SortColumnDeposit        SortColumn = "deposit"
)

func (object SortColumn) String() string {
	return string(object)
}

func SortColumnValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(SortColumn); ok {
		return enum.SortColumnValid()
	}

	return false
}

func (object SortColumn) SortColumnValid() bool {
	switch object {
	case SortColumnId, SortColumnSymbol, SortColumnInterval, SortColumnTradeDirection, SortColumnWindow, SortColumnStatus, SortColumnBind,
		SortColumnPercentIn, SortColumnPercentOut, SortColumnStopTime, SortColumnStopPercent, SortColumnDeposit:
		return true
	default:
		return false
	}
}

func (object SortColumn) DB() string {
	switch object {
	case SortColumnTradeDirection:
		return "trade_direction"
	case SortColumnBind:
		return "current_param_bind"
	case SortColumnPercentIn:
		return "current_param_percent_in"
	case SortColumnPercentOut:
		return "current_param_percent_out"
	case SortColumnStopTime:
		return "current_param_stop_time"
	case SortColumnStopPercent:
		return "current_param_stop_percent"
	default:
		return string(object)
	}
}
