package services_interface_symbol

import (
	enums_symbol "backend/internal/enums/symbol"
	models_symbol "backend/internal/models/symbol"
	"github.com/adshao/go-binance/v2/futures"
)

type SymbolService interface {
	Load(string, enums_symbol.SymbolStatus) (*models_symbol.SymbolModel, error)
	LoadAll() (*[]models_symbol.SymbolModel, error)
	LoadGroups() ([]string, error)
	Download([]futures.Symbol, []*futures.PositionRisk) error
	UpdateStatistic([]*futures.WsMarketTickerEvent) error
	Search(models_symbol.SearchRequestModel) ([]string, error)
}
