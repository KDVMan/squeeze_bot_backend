package routes_symbol

import (
	models_symbol "backend/internal/models/symbol"
	"github.com/go-chi/render"
	"net/http"
)

func (object *symbolRouteImplementation) search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_symbol.SearchRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		symbols, err := object.symbolService().Search(request)
		if err != nil {
			message := "failed to load symbolModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, symbols)
	}
}
