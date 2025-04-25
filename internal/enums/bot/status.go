package enums_bot

import "github.com/go-playground/validator/v10"

type Status string

const (
	StatusNew    Status = "new" // добавлен через калькулятор
	StatusAdd    Status = "add" // добавлен в ручную
	StatusRun    Status = "run" // запущен (собраны свечи и подписан на монету)
	StatusStop   Status = "stop"
	StatusDelete Status = "delete"
	StatusWait   Status = "wait" // защита, когда монеты льются
)

func (object Status) String() string {
	return string(object)
}

func StatusValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(Status); ok {
		return enum.StatusValid()
	}

	return false
}

func (object Status) StatusValid() bool {
	switch object {
	case StatusNew, StatusAdd, StatusRun, StatusStop, StatusDelete, StatusWait:
		return true
	default:
		return false
	}
}
