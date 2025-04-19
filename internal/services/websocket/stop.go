package services_websocket

func (object *websocketServiceImplementation) Stop() {
	object.loggerService().Info().Printf("stopping websocket service")
}
