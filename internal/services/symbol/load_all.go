package services_symbol

import (
	enums_symbol "backend/internal/enums/symbol"
	models_symbol "backend/internal/models/symbol"
)

func (object *symbolServiceImplementation) LoadAll() (*[]models_symbol.SymbolModel, error) {
	var symbols []models_symbol.SymbolModel

	if err := object.storageService().DB().
		Model(&models_symbol.SymbolModel{}).
		Where("status = ?", enums_symbol.SymbolStatusActive).
		Order("symbol").
		Find(&symbols).
		Error; err != nil {
		return nil, err
	}

	return &symbols, nil
}
