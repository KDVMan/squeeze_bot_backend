package services_interface_user

type UserService interface {
	Init()
	UpdateBalance(float64)
	UpdateAvailableBalance()
	Broadcast()
	GetAvailableBalance() float64
}
