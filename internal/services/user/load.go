package services_user

import (
	"backend/internal/enums"
	models_channel "backend/internal/models/channel"
	"errors"
	"gorm.io/gorm"
)

func (object *userServiceImplementation) Load() {
	object.userMutex.Lock()
	defer object.userMutex.Unlock()

	balance, err := object.exchangeService().UserBalance()
	if err != nil {
		object.loggerService().Error().Printf("failed to load user balance: %v", err)
		return
	}

	hedge, err := object.exchangeService().UserHedge()
	if err != nil {
		object.loggerService().Error().Printf("failed to load user hedge: %v", err)
		return
	}

	if err = object.storageService().DB().First(object.userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			object.userModel.Balance = balance
			object.userModel.AvailableBalance = balance - object.orderService().GetAmount()
			object.userModel.Hedge = hedge

			if err = object.storageService().DB().Create(object.userModel).Error; err != nil {
				object.loggerService().Error().Printf("failed to create user: %v", err)
				return
			}
		} else {
			object.loggerService().Error().Printf("failed to load user: %v", err)
			return
		}
	} else {
		object.userModel.Balance = balance
		object.userModel.AvailableBalance = balance - object.orderService().GetAmount()
		object.userModel.Hedge = hedge

		if err = object.storageService().DB().Save(object.userModel).Error; err != nil {
			object.loggerService().Error().Printf("failed to update user: %v", err)
			return
		}
	}

	object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
		Event: enums.WebsocketEventUser,
		Data:  object.userModel,
	}
}
