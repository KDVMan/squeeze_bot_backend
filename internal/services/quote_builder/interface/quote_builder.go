package services_interface_quote_builder

import models_quote "backend/internal/models/quote"

type QuoteBuilderService interface {
	Build(*models_quote.QuoteModel) *models_quote.QuoteModel
}
