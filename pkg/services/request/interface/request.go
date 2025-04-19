package services_interface_request

import "net/http"

type RequestService interface {
	Decode(http.ResponseWriter, *http.Request, interface{}) error
	Validate(http.ResponseWriter, *http.Request, interface{}) error
}
