package routes_interface_symbol_list

import "github.com/go-chi/chi/v5"

type SymbolListRoute interface {
	GetRouter() *chi.Mux
}
