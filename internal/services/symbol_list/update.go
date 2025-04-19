package services_symbol_list

import (
	models_symbol_list "backend/internal/models/symbol_list"
)

func (object *symbolListServiceImplementation) Update(request *models_symbol_list.UpdateRequestModel) (*models_symbol_list.SymbolListModel, error) {
	symbolListModel, err := object.Load()
	if err != nil {
		return nil, err
	}

	symbolListModel.Group = request.Group
	symbolListModel.Volume = request.Volume
	symbolListModel.SortColumn = request.SortColumn
	symbolListModel.SortDirection = request.SortDirection

	if err = object.storageService().DB().Save(&symbolListModel).Error; err != nil {
		return nil, err
	}

	return symbolListModel, nil
}
