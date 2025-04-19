package models_quote

type QuoteResponseModel struct {
	Quotes   []*QuoteModel `json:"quotes"`
	TimeFrom int64         `json:"timeFrom"`
	TimeTo   int64         `json:"timeTo"`
}
