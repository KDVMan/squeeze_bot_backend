package services_interface_logger

import (
	"log"
)

type LoggerService interface {
	Info() *log.Logger
	Error() *log.Logger
}
