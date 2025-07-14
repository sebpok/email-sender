package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sebpok/email-sender/internal/api/v1/handlers"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", handlers.CheckHealthHandler)

	return r
}
