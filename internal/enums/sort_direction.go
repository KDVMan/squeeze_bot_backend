package enums

import "github.com/go-playground/validator/v10"

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

func SortDirectionValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(SortDirection); ok {
		return enum.SortDirectionValid()
	}

	return false
}

func (enum SortDirection) SortDirectionValid() bool {
	switch enum {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	default:
		return false
	}
}
