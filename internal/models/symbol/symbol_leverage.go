package models_symbol

import enums_symbol "backend/internal/enums/symbol"

type SymbolLeverageModel struct {
	Level int                       `json:"level"`
	Type  enums_symbol.LeverageType `json:"type"`
}
