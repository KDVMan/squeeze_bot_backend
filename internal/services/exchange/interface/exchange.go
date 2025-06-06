package services_interface_exchange

import (
	models_bot "backend/internal/models/bot"
	"github.com/adshao/go-binance/v2/futures"
)

type ExchangeService interface {
	UserBalance() (float64, float64, error)
	UserHedge() (bool, error)
	ExchangeInfo() ([]futures.Symbol, error)
	Kline(string, string, int64, int) ([]*futures.Kline, error)
	GetListenKey() (string, error)
	DeleteListenKey() error
	Leverage() ([]*futures.PositionRisk, error)
	AddInLimit(*models_bot.BotModel, float64, float64) error
	AddOutLimit(*models_bot.BotModel, float64, float64) error
	UpdateLimit(*models_bot.BotModel, float64, float64) error
	AddOutMarket(*models_bot.BotModel, float64) error
	CancelLimit(*models_bot.BotModel) error
}
