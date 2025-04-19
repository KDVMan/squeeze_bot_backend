package routes_chart_settings

import (
	"github.com/go-chi/render"
	"net/http"
)

func (object *chartSettingsRouteImplementation) load() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chartSettings, err := object.chartSettingsService().Load()

		if err != nil {
			var message = "failed to load chartSettings"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, chartSettings)
	}
}
