package routes_interface_bot

import "github.com/go-chi/chi/v5"

type BotRoute interface {
	GetRouter() *chi.Mux
}
