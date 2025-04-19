package enums

import (
	"github.com/go-playground/validator/v10"
)

type Bind string

const (
	BindLow   Bind = "low"
	BindHigh  Bind = "high"
	BindOpen  Bind = "open"
	BindClose Bind = "close"
	BindMhl   Bind = "mhl"
	BindMoc   Bind = "moc"
)

func (enum Bind) String() string {
	return string(enum)
}

func BindValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(Bind); ok {
		return enum.BindValid()
	}

	return false
}

func (enum Bind) BindValid() bool {
	switch enum {
	case BindLow, BindHigh, BindOpen, BindClose, BindMhl, BindMoc:
		return true
	default:
		return false
	}
}
