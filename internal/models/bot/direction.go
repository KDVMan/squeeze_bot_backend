package models_bot

import (
	"backend/internal/enums"
)

type DirectionModel struct {
	Multiplier float64
	MinKeyName enums.Bind
	MaxKeyName enums.Bind
}
