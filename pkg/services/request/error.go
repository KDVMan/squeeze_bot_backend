package services_request

import (
	models_request "backend/pkg/models/request"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func (object *requestServiceImplementation) error(errors validator.ValidationErrors) models_request.ResponseModel {
	var messages []string

	for _, err := range errors {
		switch err.ActualTag() {
		case "required":
			messages = append(messages, fmt.Sprintf("field %s is a required field", err.Field()))
		case "alphanum":
			messages = append(messages, fmt.Sprintf("field %s is not valid, must be alphanumeric", err.Field()))
		case "uppercase":
			messages = append(messages, fmt.Sprintf("field %s is not valid, must be uppercase", err.Field()))
		default:
			messages = append(messages, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return models_request.ResponseModel{
		Status: "error",
		Error:  strings.Join(messages, ", "),
	}
}
