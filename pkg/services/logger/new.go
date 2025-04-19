package services_logger

import (
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"fmt"
	"log"
	"os"
	"time"
)

type loggerServiceImplementation struct {
	configService func() services_interface_config.ConfigService
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	infoFile      *os.File
	errorFile     *os.File
	currentDate   string
}

func NewLoggerService(configService func() services_interface_config.ConfigService) services_interface_logger.LoggerService {
	service := &loggerServiceImplementation{
		configService: configService,
	}

	service.updateLoggers()

	return service
}

func (object *loggerServiceImplementation) updateLoggers() {
	var err error

	config := object.configService().GetConfig()
	currentDate := time.Now().UTC().Format("2006-01-02")

	infoFileName := fmt.Sprintf("%s-%s.log", config.Logger.InfoFileName, time.Now().UTC().Format("02-01-2006"))
	errorFileName := fmt.Sprintf("%s-%s.log", config.Logger.ErrorFileName, time.Now().UTC().Format("02-01-2006"))

	if currentDate != object.currentDate {
		object.closeFiles()

		var infoLogger, errorLogger *log.Logger
		var infoFile, errorFile *os.File

		if (config.Env == "dev" && config.Logger.UseFileOnDev) || (config.Env == "prod" && config.Logger.UseFileOnProd) {
			infoFile, err = os.OpenFile(infoFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatalf("Failed to open info log file: %v", err)
			}

			errorFile, err = os.OpenFile(errorFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatalf("Failed to open error log file: %v", err)
			}

			infoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
			errorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
		} else {
			infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
			errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
		}

		object.infoLogger = infoLogger
		object.errorLogger = errorLogger
		object.infoFile = infoFile
		object.errorFile = errorFile
		object.currentDate = currentDate
	}
}

func (object *loggerServiceImplementation) Info() *log.Logger {
	object.updateLoggers()
	return object.infoLogger
}

func (object *loggerServiceImplementation) Error() *log.Logger {
	object.updateLoggers()
	return object.errorLogger
}

func (object *loggerServiceImplementation) closeFiles() {
	if object.infoFile != nil {
		err := object.infoFile.Close()
		if err != nil {
			log.Printf("Failed to close info log file: %v", err)
			return
		}

		object.infoFile = nil
	}
	if object.errorFile != nil {
		err := object.errorFile.Close()
		if err != nil {
			log.Printf("Failed to close error log file: %v", err)
			return
		}

		object.errorFile = nil
	}
}
