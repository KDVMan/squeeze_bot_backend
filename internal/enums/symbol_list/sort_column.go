package enums_symbol_list

import "github.com/go-playground/validator/v10"

type SortColumn string

const (
	SortColumnSymbol  SortColumn = "symbol"
	SortColumnPrice   SortColumn = "price"
	SortColumnVolume  SortColumn = "volume"
	SortColumnTrades  SortColumn = "trades"
	SortColumnPercent SortColumn = "percent"
)

func SortColumnValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(SortColumn); ok {
		return enum.SortColumnValid()
	}

	return false
}

func (enum SortColumn) SortColumnValid() bool {
	switch enum {
	case SortColumnSymbol, SortColumnPrice, SortColumnVolume, SortColumnTrades, SortColumnPercent:
		return true
	default:
		return false
	}
}
