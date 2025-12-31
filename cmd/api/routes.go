package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(app.recoverPanic)

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/movies", app.listMoviesHandler)
	router.Get("/v1/healthcheck", app.healthcheckHandler)
	router.Post("/v1/movies", app.createMovieHandler)
	router.Get("/v1/movies/{id}", app.showMovieHandler)
	router.Patch("/v1/movies/{id}", app.updateMovieHandler)
	router.Delete("/v1/movies/{id}", app.deleteMovieHandler)

	return router
}
