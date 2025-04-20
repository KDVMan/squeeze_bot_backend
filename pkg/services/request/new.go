package services_request

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	enums_quote "backend/internal/enums/quote"
	enums_symbol_list "backend/internal/enums/symbol_list"
	models_symbol "backend/internal/models/symbol"
	services_logger_interface "backend/pkg/services/logger/interface"
	services_request_interface "backend/pkg/services/request/interface"
	"github.com/go-playground/validator/v10"
)

type requestServiceImplementation struct {
	loggerService func() services_logger_interface.LoggerService
	validate      *validator.Validate
}

func NewRequestService(
	loggerService func() services_logger_interface.LoggerService,
) services_request_interface.RequestService {
	object := &requestServiceImplementation{
		loggerService: loggerService,
		validate:      validator.New(),
	}

	object.init()

	return object
}

func (object *requestServiceImplementation) init() {
	if err := object.validate.RegisterValidation("symbolListSortColumn", enums_symbol_list.SortColumnValidate); err != nil {
		return
	}

	if err := object.validate.RegisterValidation("sortDirection", enums.SortDirectionValidate); err != nil {
		return
	}

	if err := object.validate.RegisterValidation("symbolFormat", models_symbol.Validate); err != nil {
		return
	}

	if err := object.validate.RegisterValidation("tradeDirection", enums.TradeDirectionValidate); err != nil {
		return
	}

	if err := object.validate.RegisterValidation("interval", enums.IntervalValidate); err != nil {
		return
	}

	if err := object.validate.RegisterValidation("bind", enums.BindValidate); err != nil {
		return
	}

	if err := object.validate.RegisterValidation("quoteType", enums_quote.QuoteTypeValidate); err != nil {
		return
	}

	if err := object.validate.RegisterValidation("botSortColumn", enums_bot.SortColumnValidate); err != nil {
		return
	}

	if err := object.validate.RegisterValidation("botStatus", enums_bot.StatusValidate); err != nil {
		return
	}
}
