package routes_bot

import (
	models_bot "backend/internal/models/bot"
	"github.com/go-chi/render"
	"net/http"
)

func (object *botRouteImplementation) add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_bot.AddRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		if err := object.botService().Add(&request); err != nil {
			message := "failed to add bot"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, nil)
	}
}
