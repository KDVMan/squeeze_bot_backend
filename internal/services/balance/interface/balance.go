package services_interface_balance

type BalanceService interface {
	InitBalance(float64)
	Reserve(uint, float64) bool
	Release(uint)
}
