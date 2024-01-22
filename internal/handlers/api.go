package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// middleware for /account, check if user is authorized
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}