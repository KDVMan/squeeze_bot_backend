package services_interface_user

type UserService interface {
	Load()
	UpdateBalance(float64)
	UpdateAvailableBalance()
}
