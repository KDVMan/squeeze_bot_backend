package services_provider

import (
	services_dump "backend/pkg/services/dump"
	services_interface_dump "backend/pkg/services/dump/interface"
)

func (object *ProviderService) DumpService() services_interface_dump.DumpService {
	if object.dumpService == nil {
		object.dumpService = services_dump.NewDumpService()
	}

	return object.dumpService
}
