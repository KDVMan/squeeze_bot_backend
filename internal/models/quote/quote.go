package models_quote

import (
	"backend/internal/enums"
	"backend/internal/models"
	services_helper "backend/pkg/services/helper"
	"github.com/adshao/go-binance/v2/futures"
	"math"
)

type QuoteModel struct {
	models.DbModel
	Hash               string            `gorm:"uniqueIndex:unique_quote_01;not null" json:"-"`
	Symbol             string            `json:"symbol"`
	Interval           enums.Interval    `json:"interval"`
	TimeOpen           int64             `gorm:"uniqueIndex:unique_quote_01;not null" json:"timeOpen"`
	TimeClose          int64             `json:"timeClose"`
	Price              float64           `json:"price"`
	PriceOpen          float64           `json:"priceOpen"`
	PriceHigh          float64           `json:"priceHigh"`
	PriceLow           float64           `json:"priceLow"`
	PriceClose         float64           `json:"priceClose"`
	VolumeLeft         float64           `json:"volumeLeft"`
	VolumeRight        float64           `json:"volumeRight"`
	VolumePrice        float64           `json:"volumePrice"`
	VolumeBuyLeft      float64           `json:"volumeBuyLeft"`
	VolumeBuyRight     float64           `json:"volumeBuyRight"`
	VolumeBuyPrice     float64           `json:"volumeBuyPrice"`
	VolumeSellLeft     float64           `json:"volumeSellLeft"`
	VolumeSellRight    float64           `json:"volumeSellRight"`
	VolumeSellPrice    float64           `json:"volumeSellPrice"`
	BodySize           float64           `json:"bodySize"`
	StickUpSize        float64           `json:"stickUpSize"`
	StickDownSize      float64           `json:"stickDownSize"`
	StickRatio         float64           `json:"stickRatio"`
	CandleSize         float64           `json:"candleSize"`
	CandleBodyRange    float64           `json:"candleBodyRange"`
	Trades             int64             `json:"trades"`
	Direction          enums.Direction   `json:"direction"`
	Percent            QuotePercentModel `gorm:"embedded;embeddedPrefix:percent_" json:"percent"`
	IsClosed           bool              `json:"isClosed"`
	IsLoad             bool              `json:"isLoad"` // индикатор что это загруженные свечи (т.е. только запустили бота), при подписке этого параметра не будет
	TimeOpenFormatted  string            `gorm:"-" json:"timeOpenFormatted"`
	TimeCloseFormatted string            `gorm:"-" json:"timeCloseFormatted"`
}

func (QuoteModel) TableName() string {
	return "quotes"
}

func KlineToQuote(hash string, symbol string, interval enums.Interval, kline *futures.Kline) *QuoteModel {
	priceOpen := services_helper.MustConvertStringToFloat64(kline.Open)
	priceHigh := services_helper.MustConvertStringToFloat64(kline.High)
	priceLow := services_helper.MustConvertStringToFloat64(kline.Low)
	priceClose := services_helper.MustConvertStringToFloat64(kline.Close)
	volumeLeft := services_helper.MustConvertStringToFloat64(kline.Volume)
	volumeRight := services_helper.MustConvertStringToFloat64(kline.QuoteAssetVolume)
	volumeBuyLeft := services_helper.MustConvertStringToFloat64(kline.TakerBuyBaseAssetVolume)
	volumeBuyRight := services_helper.MustConvertStringToFloat64(kline.TakerBuyQuoteAssetVolume)
	volumeSellLeft := volumeLeft - volumeBuyLeft
	volumeSellRight := volumeRight - volumeBuyRight
	bodySize := math.Abs(priceOpen - priceClose)
	stickUpSize := priceHigh - math.Max(priceOpen, priceClose)
	stickDownSize := math.Min(priceOpen, priceClose) - priceLow
	candleSize := math.Abs(priceHigh - priceLow)
	direction := enums.DirectionUp

	if priceOpen > priceClose {
		direction = enums.DirectionDown
	}

	volumePrice := 0.0

	if volumeLeft > 0 {
		volumePrice = services_helper.Round(volumeRight/volumeLeft, 8)
	}

	volumeBuyPrice := 0.0

	if volumeBuyLeft > 0 {
		volumeBuyPrice = services_helper.Round(volumeBuyRight/volumeBuyLeft, 8)
	}

	volumeSellPrice := 0.0

	if volumeSellLeft > 0 {
		volumeSellPrice = services_helper.Round(volumeSellRight/volumeSellLeft, 8)
	}

	stickRatio := 0.0

	if stickDownSize > 0 {
		stickRatio = stickUpSize / stickDownSize
	}

	candleBodyRange := 0.0

	if candleSize > 0 {
		candleBodyRange = bodySize / candleSize
	}

	return &QuoteModel{
		Hash:               hash,
		Symbol:             symbol,
		Interval:           interval,
		TimeOpen:           kline.OpenTime,
		TimeClose:          kline.CloseTime,
		PriceOpen:          priceOpen,
		PriceHigh:          priceHigh,
		PriceLow:           priceLow,
		PriceClose:         priceClose,
		VolumeLeft:         volumeLeft,
		VolumeRight:        volumeRight,
		VolumePrice:        volumePrice,
		VolumeBuyLeft:      volumeBuyLeft,
		VolumeBuyRight:     volumeBuyRight,
		VolumeBuyPrice:     volumeBuyPrice,
		VolumeSellLeft:     volumeSellLeft,
		VolumeSellRight:    volumeSellRight,
		VolumeSellPrice:    volumeSellPrice,
		BodySize:           bodySize,
		StickUpSize:        stickUpSize,
		StickDownSize:      stickDownSize,
		StickRatio:         stickRatio,
		CandleSize:         candleSize,
		CandleBodyRange:    candleBodyRange,
		Trades:             kline.TradeNum,
		Direction:          direction,
		Percent:            GetPercent(direction, priceOpen, priceHigh, priceLow, priceClose, 2),
		IsClosed:           true,
		TimeOpenFormatted:  services_helper.MustConvertUnixMillisecondsToString(kline.OpenTime),
		TimeCloseFormatted: services_helper.MustConvertUnixMillisecondsToString(kline.CloseTime),
	}
}

func WsKlineToQuote(hash string, symbol string, interval enums.Interval, kline futures.WsKline) *QuoteModel {
	priceOpen := services_helper.MustConvertStringToFloat64(kline.Open)
	priceHigh := services_helper.MustConvertStringToFloat64(kline.High)
	priceLow := services_helper.MustConvertStringToFloat64(kline.Low)
	priceClose := services_helper.MustConvertStringToFloat64(kline.Close)
	volumeLeft := services_helper.MustConvertStringToFloat64(kline.Volume)
	volumeRight := services_helper.MustConvertStringToFloat64(kline.QuoteVolume)
	volumeBuyLeft := services_helper.MustConvertStringToFloat64(kline.ActiveBuyVolume)
	volumeBuyRight := services_helper.MustConvertStringToFloat64(kline.ActiveBuyQuoteVolume)
	volumeSellLeft := volumeLeft - volumeBuyLeft
	volumeSellRight := volumeRight - volumeBuyRight
	bodySize := math.Abs(priceOpen - priceClose)
	stickUpSize := priceHigh - math.Max(priceOpen, priceClose)
	stickDownSize := math.Min(priceOpen, priceClose) - priceLow
	candleSize := math.Abs(priceHigh - priceLow)
	direction := enums.DirectionUp

	if priceOpen > priceClose {
		direction = enums.DirectionDown
	}

	volumePrice := 0.0

	if volumeLeft > 0 {
		volumePrice = services_helper.Round(volumeRight/volumeLeft, 8)
	}

	volumeBuyPrice := 0.0

	if volumeBuyLeft > 0 {
		volumeBuyPrice = services_helper.Round(volumeBuyRight/volumeBuyLeft, 8)
	}

	volumeSellPrice := 0.0

	if volumeSellLeft > 0 {
		volumeSellPrice = services_helper.Round(volumeSellRight/volumeSellLeft, 8)
	}

	stickRatio := 0.0

	if stickDownSize > 0 {
		stickRatio = stickUpSize / stickDownSize
	}

	candleBodyRange := 0.0

	if candleSize > 0 {
		candleBodyRange = bodySize / candleSize
	}

	return &QuoteModel{
		Hash:            hash,
		Symbol:          symbol,
		Interval:        interval,
		TimeOpen:        kline.StartTime,
		TimeClose:       kline.EndTime,
		PriceOpen:       priceOpen,
		PriceHigh:       priceHigh,
		PriceLow:        priceLow,
		PriceClose:      priceClose,
		VolumeLeft:      volumeLeft,
		VolumeRight:     volumeRight,
		VolumePrice:     volumePrice,
		VolumeBuyLeft:   volumeBuyLeft,
		VolumeBuyRight:  volumeBuyRight,
		VolumeBuyPrice:  volumeBuyPrice,
		VolumeSellLeft:  volumeSellLeft,
		VolumeSellRight: volumeSellRight,
		VolumeSellPrice: volumeSellPrice,
		BodySize:        bodySize,
		StickUpSize:     stickUpSize,
		StickDownSize:   stickDownSize,
		StickRatio:      stickRatio,
		CandleSize:      candleSize,
		CandleBodyRange: candleBodyRange,
		Trades:          kline.TradeNum,
		Direction:       direction,
		Percent:         GetPercent(direction, priceOpen, priceHigh, priceLow, priceClose, 2),
		IsClosed:        kline.IsFinal,
	}
}
