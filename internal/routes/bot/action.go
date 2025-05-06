package routes_bot

import (
	models_bot "backend/internal/models/bot"
	"github.com/go-chi/render"
	"net/http"
)

func (object *botRouteImplementation) action() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_bot.ActionRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		if err := object.botService().Action(&request); err != nil {
			message := "failed to action bot"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, nil)
	}
}
