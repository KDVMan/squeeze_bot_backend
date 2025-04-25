package models_user

type UserModel struct {
	Balance          float64 `json:"balance"`
	AvailableBalance float64 `json:"availableBalance"`
	Hedge            bool    `json:"hedge"`
}
