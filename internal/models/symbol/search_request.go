package models_symbol

type SearchRequestModel struct {
	Symbol string `json:"symbol" validate:"required,symbolFormat,uppercase"`
}
