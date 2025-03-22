package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	middleware2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/middleware"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	userService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
)

func New(
	logger zap.SugaredLogger,
	userService userService.UserService,
	rolesService roles.RolesService,
	jwtSecret string,
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
			userHandler := user.NewUserHandler(logger, userService, rolesService)

			r.Route(
				"/user", func(r chi.Router) {
					r.Post("/register", userHandler.Register())
					r.Post("/login", userHandler.Login())

					// Protected endpoints: wrap these with auth middleware.
					r.Group(func(r chi.Router) {
						r.Use(middleware2.AuthMiddleware([]byte(jwtSecret)))
						r.Post("/update-details", userHandler.UpdateDetails())
						r.Delete("/", userHandler.Delete())
					})
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
