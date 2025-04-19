package services_provider

import (
	services_request "backend/pkg/services/request"
	services_interface_request "backend/pkg/services/request/interface"
)

func (object *ProviderService) RequestService() services_interface_request.RequestService {
	if object.requestService == nil {
		object.requestService = services_request.NewRequestService(
			object.LoggerService,
		)
	}

	return object.requestService
}
