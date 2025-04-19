package models_symbol_list

import (
	"backend/internal/enums"
	enums_symbol_list "backend/internal/enums/symbol_list"
)

type SymbolListModel struct {
	ID            int                          `json:"-" gorm:"primaryKey"`
	Group         string                       `json:"group"`
	Volume        int                          `json:"volume"`
	SortColumn    enums_symbol_list.SortColumn `json:"sortColumn"`
	SortDirection enums.SortDirection          `json:"sortDirection"`
}

func (SymbolListModel) TableName() string {
	return "symbol_list"
}

func LoadDefault() *SymbolListModel {
	return &SymbolListModel{
		Group:         "USDT",
		Volume:        0,
		SortColumn:    enums_symbol_list.SortColumnVolume,
		SortDirection: enums.SortDirectionDesc,
	}
}
