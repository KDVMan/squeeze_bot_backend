package services_provider

import (
	services_quote "backend/internal/services/quote"
	services_quote_interface "backend/internal/services/quote/interface"
)

func (object *ProviderService) QuoteService() services_quote_interface.QuoteService {
	if object.quoteService == nil {
		object.quoteService = services_quote.NewQuoteService(
			object.LoggerService,
			object.StorageService,
			object.WebsocketService,
			object.DumpService,
			object.ExchangeService,
			object.ExchangeWebsocketService,
			object.InitService,
			object.BotRepositoryService,
			object.ExchangeLimitService,
		)
	}

	return object.quoteService
}
