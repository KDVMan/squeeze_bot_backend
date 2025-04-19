package services_interface_symbol_list

import (
	models_symbol_list "backend/internal/models/symbol_list"
)

type SymbolListService interface {
	Load() (*models_symbol_list.SymbolListModel, error)
	Update(*models_symbol_list.UpdateRequestModel) (*models_symbol_list.SymbolListModel, error)
}
