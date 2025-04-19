package services_storage

import (
	models_bot "backend/internal/models/bot"
	models_chart_settings "backend/internal/models/chart_settings"
	models_exchange_limit "backend/internal/models/exchange_limit"
	models_init "backend/internal/models/init"
	models_quote "backend/internal/models/quote"
	models_symbol "backend/internal/models/symbol"
	models_symbol_list "backend/internal/models/symbol_list"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type storageServiceImplementation struct {
	db *gorm.DB
}

func NewStorageService(configService func() services_interface_config.ConfigService) services_interface_storage.StorageService {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(sqlite.Open(configService().GetConfig().DB.Path), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)
	}

	if err := db.Exec("PRAGMA foreign_keys = ON", nil).Error; err != nil {
		log.Fatalf("Failed to enable foreign keys: %s", err)
	}

	err = db.AutoMigrate(
		&models_init.InitModel{},
		&models_symbol.SymbolModel{},
		&models_symbol_list.SymbolListModel{},
		&models_exchange_limit.ExchangeLimitModel{},
		&models_quote.QuoteModel{},
		&models_chart_settings.ChartSettings{},
		&models_bot.BotModel{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	return &storageServiceImplementation{
		db: db,
	}
}

func (object *storageServiceImplementation) DB() *gorm.DB {
	return object.db
}
