package services_symbol_list

import (
	models_symbol_list "backend/internal/models/symbol_list"
	"errors"
	"gorm.io/gorm"
)

func (object *symbolListServiceImplementation) Load() (*models_symbol_list.SymbolListModel, error) {
	var symbolListModel *models_symbol_list.SymbolListModel

	if err := object.storageService().DB().First(&symbolListModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			symbolListModel = models_symbol_list.LoadDefault()

			if err = object.storageService().DB().Create(symbolListModel).Error; err != nil {
				return nil, err
			}

			return symbolListModel, nil
		}

		return nil, err
	}

	return symbolListModel, nil
}
