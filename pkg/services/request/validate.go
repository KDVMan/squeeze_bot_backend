package services_request

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (object *requestServiceImplementation) Validate(w http.ResponseWriter, r *http.Request, request interface{}) error {
	if err := object.validate.Struct(request); err != nil {
		var validateError validator.ValidationErrors

		if errors.As(err, &validateError) {
			object.loggerService().Error().Printf("%s: %v", "validation failed", err)

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, object.error(validateError))

			return err
		}

		message := "unknown validation error"
		object.loggerService().Error().Printf("%s: %v", message, err)

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, message)

		return err
	}

	return nil
}
