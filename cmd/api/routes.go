package main

import (
	"expvar"
	"greenlight/internal/ui"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	// Convert the notFoundResponse() helper to a http.Handler using the
	// http.HandlerFunc() adapter, and then set it as the custom error handler for 404
	// Not Found responses.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// Likewise, convert the methodNotAllowedResponse() helper to a http.Handler and set
	// it as the custom error handler for 405 Method Not Allowed responses.
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/movies", app.requirePermission("movies:read", app.listMoviesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.requirePermission("movies:write", app.createMovieHandler))
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requirePermission("movies:read", app.showMovieHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.requirePermission("movies:write", app.updateMovieHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.requirePermission("movies:write", app.deleteMovieHandler))

	// set requirements
	router.HandlerFunc(http.MethodGet, "/v1/users", app.getAllUsersHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/users", app.deleteUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	// static file server
	router.Handler(http.MethodGet, "/templates/static/", http.FileServer(http.Dir("/static")))

	router.Handler(http.MethodGet, "/debug/vars", app.metrics(expvar.Handler()))
	// tests
	router.HandlerFunc(http.MethodPost, "/test/readJSON", app.testReadJSON)
	router.HandlerFunc(http.MethodGet, "/test/home", ui.HTML("home"))
	router.HandlerFunc(http.MethodGet, "/test/registration", ui.HTML("registration"))

	return app.recoverPanic(app.enableCORS(app.rateLimit3(app.authenticate(router))))
}
