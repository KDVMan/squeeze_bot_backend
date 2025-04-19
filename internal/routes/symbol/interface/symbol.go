package routes_interface_symbol

import "github.com/go-chi/chi/v5"

type SymbolRoute interface {
	GetRouter() *chi.Mux
}
