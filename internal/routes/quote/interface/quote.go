package routes_interface_quote

import "github.com/go-chi/chi/v5"

type QuoteRoute interface {
	GetRouter() *chi.Mux
}
