package services_symbol

import (
	models_symbol "backend/internal/models/symbol"
)

func (object *symbolServiceImplementation) LoadGroups() ([]string, error) {
	var groups []string

	if err := object.storageService().DB().
		Model(&models_symbol.SymbolModel{}).
		Distinct("`group`").
		Order("`group` ASC").
		Pluck("`group`", &groups).Error; err != nil {
		return nil, err
	}

	return groups, nil
}
