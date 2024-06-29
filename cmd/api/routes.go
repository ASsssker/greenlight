package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewMux()

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Route("/v1/movies", func(r chi.Router) {
		r.Post("/", app.createMovieHandler)
		r.Get("/{id}", app.showMovieHandler)
	})

	return router
}
