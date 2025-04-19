package routes_symbol_list

import (
	"github.com/go-chi/render"
	"net/http"
)

func (object *symbolListRouteImplementation) load() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbolListModel, err := object.symbolListService().Load()
		if err != nil {
			var message = "failed to load symbolListModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, symbolListModel)
	}
}
