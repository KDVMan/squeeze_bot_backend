package routes_quote

import (
	models_quote "backend/internal/models/quote"
	"github.com/go-chi/render"
	"net/http"
)

func (object *quoteRouteImplementation) load() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_quote.QuoteRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		quotes, err := object.quoteService().Load(&request)

		if err != nil {
			message := "failed to load quotes"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, quotes)
	}
}
