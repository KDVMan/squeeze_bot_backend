package routes_init

import (
	"github.com/go-chi/render"
	"net/http"
)

func (object *initRouteImplementation) load() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		initModel, err := object.initService().Load()
		if err != nil {
			var message = "failed to load initModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, initModel)
	}
}
