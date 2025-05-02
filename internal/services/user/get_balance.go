package services_user

func (object *userServiceImplementation) GetBalance() float64 {
	object.userMutex.Lock()
	defer object.userMutex.Unlock()

	return object.userModel.Balance
}
