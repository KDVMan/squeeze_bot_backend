package enums_quote

import "github.com/go-playground/validator/v10"

type QuoteType string

const (
	QuoteTypeInit QuoteType = "init"
	QuoteTypeLoad QuoteType = "load"
)

func QuoteTypeValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(QuoteType); ok {
		return enum.QuoteTypeValid()
	}

	return false
}

func (enum QuoteType) QuoteTypeValid() bool {
	switch enum {
	case QuoteTypeInit, QuoteTypeLoad:
		return true
	default:
		return false
	}
}
