package services_init

import (
	services_interface_init "backend/internal/services/init/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type initServiceImplementation struct {
	storageService   func() services_interface_storage.StorageService
	websocketService func() services_interface_websocket.WebsocketService
	symbolService    func() services_interface_symbol.SymbolService
	// calculatorService       func() services_interface_calculator.CalculatorService
	// calculatorPresetService func() services_calculator_preset_interface.CalculatorPresetService
}

func NewInitService(
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	symbolService func() services_interface_symbol.SymbolService,
	// calculatorService func() services_interface_calculator.CalculatorService,
	// calculatorPresetService func() services_calculator_preset_interface.CalculatorPresetService,
) services_interface_init.InitService {
	return &initServiceImplementation{
		storageService:   storageService,
		websocketService: websocketService,
		symbolService:    symbolService,
		// calculatorService:       calculatorService,
		// calculatorPresetService: calculatorPresetService,
	}
}
