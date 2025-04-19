package services_init

import (
	models_init "backend/internal/models/init"
	"errors"
	"gorm.io/gorm"
)

func (object *initServiceImplementation) Load() (*models_init.InitModel, error) {
	var initModel *models_init.InitModel

	if err := object.storageService().DB().First(&initModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			initModel = models_init.LoadDefault()

			if err = object.storageService().DB().Create(initModel).Error; err != nil {
				return nil, err
			}

			return initModel, nil
		}

		return nil, err
	}

	return initModel, nil
}
