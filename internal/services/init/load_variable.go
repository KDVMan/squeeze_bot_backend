package services_init

import (
	models_init "backend/internal/models/init"
)

func (object *initServiceImplementation) LoadVariable() (*models_init.InitVariableModel, error) {
	var initVariableModel = &models_init.InitVariableModel{
		Groups: []string{},
	}

	groups, err := object.symbolService().LoadGroups()
	if err == nil {
		initVariableModel.Groups = groups
	}

	return initVariableModel, nil
}
