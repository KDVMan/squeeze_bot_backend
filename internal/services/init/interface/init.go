package services_interface_init

import models_init "backend/internal/models/init"

type InitService interface {
	Load() (*models_init.InitModel, error)
	LoadVariable() (*models_init.InitVariableModel, error)
	Update(*models_init.UpdateRequestModel) (*models_init.InitModel, error)
}
