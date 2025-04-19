package routes_chart_settings

import (
	models_chart_settings "backend/internal/models/chart_settings"
	"github.com/go-chi/render"
	"net/http"
)

func (object *chartSettingsRouteImplementation) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models_chart_settings.UpdateRequestModel

		if err := object.requestService().Decode(w, r, &request); err != nil {
			return
		}

		if err := object.requestService().Validate(w, r, &request); err != nil {
			return
		}

		chartSettingsModel, err := object.chartSettingsService().Update(&request)

		if err != nil {
			message := "failed to update chartSettingsModel"
			object.loggerService().Error().Printf("%s: %v", message, err)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, message)

			return
		}

		render.JSON(w, r, chartSettingsModel)
	}
}
