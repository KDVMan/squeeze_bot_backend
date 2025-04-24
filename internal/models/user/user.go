package models_user

type UserModel struct {
	ID               int     `json:"-" gorm:"primaryKey"`
	Balance          float64 `json:"balance"`
	AvailableBalance float64 `json:"availableBalance"`
	Hedge            bool    `json:"hedge"`
}

func (UserModel) TableName() string {
	return "user"
}
