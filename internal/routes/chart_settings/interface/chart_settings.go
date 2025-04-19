package routes_interface_chart_settings

import "github.com/go-chi/chi/v5"

type ChartSettingsRoute interface {
	GetRouter() *chi.Mux
}
