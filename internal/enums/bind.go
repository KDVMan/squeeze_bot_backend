package enums

import (
	"github.com/go-playground/validator/v10"
	"math/rand"
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

func BindValues() []Bind {
	return []Bind{
		BindLow, BindHigh, BindOpen, BindClose, BindMhl, BindMoc,
	}
}

func BindRandom(binds []Bind) Bind {
	if len(binds) == 0 {
		return ""
	}

	return binds[rand.Intn(len(binds))]
}

func BindValidate(field validator.FieldLevel) bool {
	binds, ok := field.Field().Interface().([]Bind)

	if !ok {
		return false
	}

	validBinds := BindValues()

	for _, bind := range binds {
		if !bindContains(validBinds, bind) {
			return false
		}
	}
	return true
}

func bindContains(s []Bind, e Bind) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
