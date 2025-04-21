package services_provider

import (
	services_quote_repository "backend/internal/services/quote_repository"
	services_interface_quote_repository "backend/internal/services/quote_repository/interface"
)

func (object *ProviderService) QuoteRepositoryService() services_interface_quote_repository.QuoteRepositoryService {
	if object.quoteRepositoryService == nil {
		object.quoteRepositoryService = services_quote_repository.NewQuoteRepositoryService(
			object.LoggerService,
		)
	}

	return object.quoteRepositoryService
}
