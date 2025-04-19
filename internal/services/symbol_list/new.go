package services_symbol_list

import (
	services_interface_symbol_list "backend/internal/services/symbol_list/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type symbolListServiceImplementation struct {
	storageService func() services_interface_storage.StorageService
}

func NewSymbolListService(
	storageService func() services_interface_storage.StorageService,
) services_interface_symbol_list.SymbolListService {
	return &symbolListServiceImplementation{
		storageService: storageService,
	}
}
