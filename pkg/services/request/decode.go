package services_request

import (
	"github.com/go-chi/render"
	"net/http"
)

func (object *requestServiceImplementation) Decode(w http.ResponseWriter, r *http.Request, request interface{}) error {
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		message := "failed to decode request body"
		object.loggerService().Error().Printf("%s: %v", message, err)

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, message)

		return err
	}

	return nil
}
