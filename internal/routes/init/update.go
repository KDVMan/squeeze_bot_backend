package routes_init

import (
	models_init "backend/internal/models/init"
	"github.com/go-chi/render"
	"net/http"
)

func (object *initRouteImplementation) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_init.UpdateRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		initModel, err := object.initService().Update(&request)
		if err != nil {
			message := "failed to update initModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, initModel)
	}
}
