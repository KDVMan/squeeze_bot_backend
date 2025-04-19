package models_symbol

import (
	enums_symbol "backend/internal/enums/symbol"
	"backend/internal/models"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type SymbolModel struct {
	models.DbModel
	Group     string                    `gorm:"not null,index" json:"group"`
	Symbol    string                    `gorm:"unique;not null" json:"symbol"`
	Status    enums_symbol.SymbolStatus `gorm:"not null" json:"status"`
	Leverage  SymbolLeverageModel       `gorm:"embedded;embeddedPrefix:leverage_" json:"leverage"`
	Limit     SymbolLimitModel          `gorm:"embedded;embeddedPrefix:limit_" json:"limit"`
	Statistic SymbolStatisticModel      `gorm:"embedded;embeddedPrefix:statistic_" json:"statistic"`
}

func (SymbolModel) TableName() string {
	return "symbols"
}

func Validate(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[A-Z0-9_]+$`)
	return re.MatchString(fl.Field().String())
}
