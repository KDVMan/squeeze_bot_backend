package services_interface_user

import (
	models_user "backend/internal/models/user"
)

type UserService interface {
	Load()
	Update(*models_user.UserModel, bool)
	UpdateBalance(float64)
}
