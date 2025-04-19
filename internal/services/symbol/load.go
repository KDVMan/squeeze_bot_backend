package services_symbol

import (
	enums_symbol "backend/internal/enums/symbol"
	models_symbol "backend/internal/models/symbol"
)

func (object *symbolServiceImplementation) Load(symbol string, status enums_symbol.SymbolStatus) (*models_symbol.SymbolModel, error) {
	var symbolModel models_symbol.SymbolModel

	if err := object.storageService().DB().Where("symbol = ? AND status = ?", symbol, status).First(&symbolModel).Error; err != nil {
		return nil, err
	}

	return &symbolModel, nil
}
