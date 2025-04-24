package services_symbol

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
	models_symbol "backend/internal/models/symbol"
	service_helper "backend/pkg/services/helper"
	"errors"
	"github.com/adshao/go-binance/v2/futures"
	"gorm.io/gorm"
)

func (object *symbolServiceImplementation) UpdateStatistic(tickets []*futures.WsMarketTickerEvent) error {
	err := object.storageService().DB().Transaction(func(tx *gorm.DB) error {
		for _, ticket := range tickets {
			var symbolModel models_symbol.SymbolModel

			if err := tx.Where("symbol = ?", ticket.Symbol).First(&symbolModel).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					continue
				}

				return err
			}

			symbolModel.Statistic = models_symbol.SymbolStatisticModel{
				Price:        service_helper.MustConvertStringToFloat64(ticket.ClosePrice),
				PriceLow:     service_helper.MustConvertStringToFloat64(ticket.LowPrice),
				PriceHigh:    service_helper.MustConvertStringToFloat64(ticket.HighPrice),
				PricePercent: service_helper.MustConvertStringToFloat64(ticket.PriceChangePercent),
				Volume:       service_helper.MustConvertStringToFloat64(ticket.QuoteVolume),
				Trades:       ticket.TradeCount,
			}

			if err := tx.Save(&symbolModel).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err == nil {
		symbols, err := object.LoadAll()
		if err != nil {
			return err
		}

		broadcastModel := models_channel.BroadcastChannelModel{
			Event: enums.WebsocketEventSymbolList,
			Data:  symbols,
		}

		object.websocketService().GetBroadcastChannel() <- &broadcastModel
	}

	return err
}
