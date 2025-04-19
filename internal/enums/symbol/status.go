package enums_symbol

type SymbolStatus string

const (
	SymbolStatusActive   SymbolStatus = "active"
	SymbolStatusInactive SymbolStatus = "inactive"
	SymbolStatusUnknown  SymbolStatus = "unknown"
)
