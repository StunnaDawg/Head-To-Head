package handlers

import (
	"github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
	"github.com/StunnaDawg/Head-To-Head/internal/middlewares"
)

func Handler(r *chi.Mux) {
	r.Use(middleware.StripSlashes)

	r.Route("/chosen", func(route, chi.Router) {
		router.Use(middlewares.Authorization)
		router.Get("/athlete", GetFavouriteAthlete)
	})
}