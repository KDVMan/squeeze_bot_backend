package services_bot

import (
	models_bot "backend/internal/models/bot"
	"errors"
	"gorm.io/gorm"
)

func (object *botServiceImplementation) LoadByHash(hash string) *models_bot.BotModel {
	var botModel models_bot.BotModel

	if err := object.storageService().DB().
		Where("hash = ?", hash).
		First(&botModel).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			object.loggerService().Error().Printf("failed to load bot by hash: %v", err)
		}

		return nil
	}

	return &botModel
}
