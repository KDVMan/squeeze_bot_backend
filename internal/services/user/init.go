package services_user

import (
	models_user "backend/internal/models/user"
)

func (object *userServiceImplementation) Init() {
	object.userMutex.Lock()
	defer object.userMutex.Unlock()

	balance, availableBalance, err := object.exchangeService().UserBalance()
	if err != nil {
		object.loggerService().Error().Printf("failed to load user balance: %v", err)
		return
	}

	hedge := true

	// hedge, err := object.exchangeService().UserHedge()
	// if err != nil {
	// 	object.loggerService().Error().Printf("failed to load user hedge: %v", err)
	// 	return
	// }

	object.userModel = &models_user.UserModel{
		Balance:          balance,
		AvailableBalance: availableBalance,
		Hedge:            hedge,
	}
}
