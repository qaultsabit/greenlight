package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.NotFound(app.notFoundResponse)
	mux.MethodNotAllowed(app.methodNotAllowedResponse)

	mux.Get("/v1/healthcheck", app.healthcheckHandler)
	mux.Post("/v1/movies", app.createMovieHandler)
	mux.Get("/v1/movies/{id}", app.showMovieHandler)
	mux.Patch("/v1/movies/{id}", app.updateMovieHandler)
	mux.Delete("/v1/movies/{id}", app.deleteMovieHandler)

	return mux
}
