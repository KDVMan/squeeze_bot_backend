package services_user

func (object *userServiceImplementation) GetAvailableBalance() float64 {
	object.userMutex.Lock()
	defer object.userMutex.Unlock()

	return object.userModel.AvailableBalance
}
