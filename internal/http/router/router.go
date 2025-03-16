package router

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func New() http.Handler {
	// Create a new Chi router.
	router := chi.NewRouter()

	// Add middleware.
	router.Use(middleware.Logger)    // logs every request
	router.Use(middleware.Recoverer) // recovers from panics

	// Define the /alive endpoint.
	registerAliveEndpoint(router)
	router.Route(
		"/api/v1", func(r chi.Router) {

			//h := authhandler.NewHandler(authService, logger)

			r.Route(
				"/auth", func(r chi.Router) {
					//r.Use(authMiddleware)
					//r.Get("/{id}", h.Read())
				},
			)
		},
	)

	return router
}

func registerAliveEndpoint(router *chi.Mux) {
	router.Get("/alive", func(w http.ResponseWriter, r *http.Request) {
		// Return a simple status message.
		render.Json(w, http.StatusOK, "API is alive!")
	})
}
