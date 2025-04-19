package enums

import "github.com/go-playground/validator/v10"

type TradeDirection string

const (
	TradeDirectionLong  TradeDirection = "long"
	TradeDirectionShort TradeDirection = "short"
)

func TradeDirectionValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(TradeDirection); ok {
		return enum.TradeDirectionValid()
	}

	return false
}

func (enum TradeDirection) TradeDirectionValid() bool {
	switch enum {
	case TradeDirectionLong, TradeDirectionShort:
		return true
	default:
		return false
	}
}
