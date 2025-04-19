package routes_symbol_list

import (
	models_symbol_list "backend/internal/models/symbol_list"
	"github.com/go-chi/render"
	"net/http"
)

func (object *symbolListRouteImplementation) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_symbol_list.UpdateRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		symbolListModel, err := object.symbolListService().Update(&request)

		if err != nil {
			message := "failed to update symbolListModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, symbolListModel)
	}
}
