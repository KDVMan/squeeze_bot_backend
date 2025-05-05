package services_balance

func (object *balanceServiceImplementation) Reserve(botID uint, amount float64) bool {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	if object.available() < amount {
		return false
	}

	object.data[botID] = amount

	return true
}

func (object *balanceServiceImplementation) Release(botID uint) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	delete(object.data, botID)
}

func (object *balanceServiceImplementation) available() float64 {
	totalAmount := 0.0

	for _, amount := range object.data {
		totalAmount += amount
	}

	return object.balance - totalAmount
}

func (object *balanceServiceImplementation) UpdateBalance(delta float64) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	object.balance += delta
}
