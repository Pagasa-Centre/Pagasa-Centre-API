package router

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth"
	authentication "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
)

func New(
	logger zap.SugaredLogger,
	authService authentication.Service,
) http.Handler {
	// Create a new Chi router.
	router := chi.NewRouter()

	// Add middleware.
	router.Use(middleware.Logger)    // logs every request
	router.Use(middleware.Recoverer) // recovers from panics

	// Define the /alive endpoint.
	registerAliveEndpoint(router)
	router.Route(
		"/api/v1", func(r chi.Router) {

			authHandler := auth.NewHandler(logger, authService)

			r.Route(
				"/auth", func(r chi.Router) {
					//r.Use(authMiddleware)
					r.Post("/register", authHandler.Register())
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
