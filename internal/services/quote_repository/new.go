package services_quote_repository

import (
	models_quote "backend/internal/models/quote"
	services_interface_quote_repository "backend/internal/services/quote_repository/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
)

type quoteRepositoryServiceImplementation struct {
	loggerService func() services_interface_logger.LoggerService
	data          map[string][]*models_quote.QuoteModel
	mutex         *sync.Mutex
}

func NewQuoteRepositoryService(
	loggerService func() services_interface_logger.LoggerService,
) services_interface_quote_repository.QuoteRepositoryService {
	return &quoteRepositoryServiceImplementation{
		loggerService: loggerService,
		data:          make(map[string][]*models_quote.QuoteModel),
		mutex:         &sync.Mutex{},
	}
}
