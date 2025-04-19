package services_symbol

import (
	enums_symbol "backend/internal/enums/symbol"
	models_symbol "backend/internal/models/symbol"
	"fmt"
)

func (object *symbolServiceImplementation) Search(request models_symbol.SearchRequestModel) ([]string, error) {
	var symbols []string

	query := fmt.Sprintf("%%%s%%", request.Symbol)

	err := object.storageService().DB().
		Model(&models_symbol.SymbolModel{}).
		Where("UPPER(symbol) LIKE UPPER(?) AND status = ?", query, enums_symbol.SymbolStatusActive).
		Order("symbol").
		Pluck("symbol", &symbols).
		Error

	if err != nil {
		return nil, err
	}

	return symbols, nil
}
