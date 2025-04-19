package enums_symbol

import (
	"strings"
)

type LeverageType string

const (
	LeverageTypeIsolated LeverageType = "isolated"
	LeverageTypeCrossed  LeverageType = "cross"
	LeverageTypeUnknown  LeverageType = "unknown"
)

func ConvertLeverage(value string) LeverageType {
	switch strings.ToLower(value) {
	case string(LeverageTypeIsolated):
		return LeverageTypeIsolated
	case string(LeverageTypeCrossed):
		return LeverageTypeCrossed
	default:
		return LeverageTypeUnknown
	}
}
