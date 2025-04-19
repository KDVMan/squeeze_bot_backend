package models_user

type UserModel struct {
	TotalBalance     float64 `json:"totalBalance"`
	AvailableBalance float64 `json:"availableBalance"`
	Hedge            bool    `json:"hedge"`
}

func LoadDefault() *UserModel {
	return &UserModel{
		TotalBalance:     0,
		AvailableBalance: 0,
		Hedge:            false,
	}
}
