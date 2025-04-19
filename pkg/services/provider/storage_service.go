package services_provider

import (
	services_storage "backend/pkg/services/storage"
	services_interface_storage "backend/pkg/services/storage/interface"
)

func (object *ProviderService) StorageService() services_interface_storage.StorageService {
	if object.storageService == nil {
		object.storageService = services_storage.NewStorageService(
			object.ConfigService,
		)
	}

	return object.storageService
}
