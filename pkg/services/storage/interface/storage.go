package services_interface_storage

import (
	"gorm.io/gorm"
)

type StorageService interface {
	DB() *gorm.DB
}
