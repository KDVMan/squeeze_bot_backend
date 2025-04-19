package routes_init

import (
	"github.com/go-chi/render"
	"net/http"
)

func (object *initRouteImplementation) loadVariable() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		initVariableModel, err := object.initService().LoadVariable()

		if err != nil {
			var message = "failed to load initVariableModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, initVariableModel)
	}
}
