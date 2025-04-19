package routes_interface_init

import "github.com/go-chi/chi/v5"

type InitRoute interface {
	GetRouter() *chi.Mux
}
