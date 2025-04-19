package models_init

import (
	enums_symbol "backend/internal/enums/symbol"
	models_quote "backend/internal/models/quote"
	"encoding/json"
	"gorm.io/gorm"
)

type InitModel struct {
	ID            int                       `json:"-" gorm:"primaryKey"`
	Symbol        string                    `json:"symbol"`
	Intervals     []*models_quote.Interval  `json:"intervals" gorm:"-"`
	IntervalsJson string                    `json:"-"`
	QuotesLimit   uint                      `json:"quotesLimit"`
	Precision     int                       `json:"precision"`
	LeverageLevel int                       `json:"leverageLevel"`
	LeverageType  enums_symbol.LeverageType `json:"leverageType"`
}

func (InitModel) TableName() string {
	return "init"
}

func LoadDefault() *InitModel {
	return &InitModel{
		Symbol:        "BTCUSDT",
		Intervals:     models_quote.IntervalLoadDefault(),
		QuotesLimit:   1500,
		Precision:     2,
		LeverageLevel: 0,
		LeverageType:  enums_symbol.LeverageTypeUnknown,
	}
}

func (object *InitModel) BeforeSave(tx *gorm.DB) (err error) {
	if object.Intervals != nil {
		data, err := json.Marshal(object.Intervals)
		if err != nil {
			return err
		}

		object.IntervalsJson = string(data)
	}

	return nil
}

func (object *InitModel) AfterFind(tx *gorm.DB) (err error) {
	if object.IntervalsJson != "" {
		var list []*models_quote.Interval

		err = json.Unmarshal([]byte(object.IntervalsJson), &list)
		if err != nil {
			return err
		}

		object.Intervals = list
	}

	return nil
}
