package services_bot

import (
	enums_bot "backend/internal/enums/bot"
)

func (object *botServiceImplementation) GetAmount() float64 {
	amount := 0.0

	for _, botModel := range object.botRepositoryService().GetAll() {
		if botModel.Deal.Status == enums_bot.DealStatusSendOpenLimit ||
			botModel.Deal.Status == enums_bot.DealStatusOpenLimit ||
			botModel.Deal.Status == enums_bot.DealStatusOpen {
			amount += botModel.Deposit
		}
	}

	return amount
}
