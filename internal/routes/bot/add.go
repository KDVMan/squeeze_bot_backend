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
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]any{"success": false, "error": "invalid request"})
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]any{"success": false, "error": "validation failed"})
			return
		}

		if err := object.botService().Add(&request); err != nil {
			message := "failed to add bot"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]any{"success": false, "error": message})

			return
		}

		render.JSON(w, r, map[string]bool{"success": true})
	}
}
