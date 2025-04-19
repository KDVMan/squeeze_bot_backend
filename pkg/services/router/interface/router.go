package services_interface_router

import (
	"github.com/go-chi/chi/v5"
)

type RouterService interface {
	GetRouter() *chi.Mux
}
