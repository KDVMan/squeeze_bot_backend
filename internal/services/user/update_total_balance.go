package services_user

func (object *userServiceImplementation) UpdateBalance(balance float64) {
	object.userModel.TotalBalance += balance
	object.userModel.AvailableBalance += balance

	object.Update(object.userModel, true)
}
